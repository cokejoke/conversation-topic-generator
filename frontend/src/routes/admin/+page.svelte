<script lang="ts">

import PackComponent from "$lib/components/PackComponent.svelte";
import {type Pack} from "../../lib/model/pack";
import {axiosInstance} from "../../lib/variables";
import {onMount} from "svelte";

let packs: Array<Pack>;
async function load() {
    const response = await axiosInstance.get('/pack');
    packs = response.data;
}

onMount(async () => load());

</script>

<svelte:head>
    <title>Admin | Gesprächsthema Generator</title>
</svelte:head>

<div class="container">
    <div class="row">
        {#if packs}
            {#each packs as pack}
                <PackComponent pack="{pack}" />
            {/each}
        {/if}
    </div>
</div>

<style>

    .container {
        padding: 90px 0;
        width: 100%;
        max-width: 1500px;
        margin: 0 auto;
    }

    .row {
        display: grid;
        grid-template-columns: repeat(4, 1fr);
        column-gap: 40px;
    }

    @media all and (max-device-width: 1420px){
        .row {
            grid-template-columns: repeat(3, 1fr);
        }
    }

    @media all and (max-device-width: 1200px){
        .row {
            grid-template-columns: repeat(2, 1fr);
        }
    }

    @media all and (max-device-width: 850px){
        .row {
            grid-template-columns: auto;
        }
    }
</style>