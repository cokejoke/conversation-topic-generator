<script lang="ts">
    /** @type {import('./$types').PageData} */
    import {onMount} from "svelte";
    import {type Pack} from "$lib/model/pack";
    import {axiosInstance} from "$lib/variables";
    import { fade } from 'svelte/transition';
    import TopicComponent from "$lib/components/TopicComponent.svelte";

    export let data;
    let id: number = data.id;

    let pack: Pack;

    onMount(async () => {
        const response = await axiosInstance.get(`/pack/${id}`);
        pack = response.data;
    });
</script>

{#if pack}
    <div class="container" in:fade>
        <div class="header">
            <div class="info">
                <img src="{pack.imageUrl ? pack.imageUrl : '/placeholder.png'}" alt="Pack image">
                <h3>{pack.name}</h3>
            </div>
            <div class="buttons">
                <a href="/admin/pack/delete/{id}"><img src="/svg/delete.svg" alt="Delete"></a>
                <a href="/admin/topic/add/{id}"><img src="/svg/add_box.svg" alt="Add"></a>
            </div>
        </div>
        <div class="single-buttons">
            <a href="/admin/pack/delete/{id}"><img src="/svg/delete.svg" alt="-">Löschen</a>
            <a href="/admin/topic/add/{id}"><img src="/svg/add_box.svg" alt="+">Hinzufügen</a>
        </div>
        <div class="topics">
            {#each pack.topics as topic}
                <TopicComponent topic="{topic}" />
            {/each}
        </div>
    </div>
{/if}

<style>

    .container {
        padding: 90px 0;
        width: 100%;
        max-width: 1500px;
        margin: 0 auto;
    }

    .header {
        display: grid;
        grid-template-columns: 3fr 1fr;
        padding: 15px 50px;
        background: #DF7599;
        align-items: center;
        border-radius: 20px;
    }

    .info {
        display: grid;
        grid-gap: 20px;
        align-items: center;
        grid-template-columns: 65px auto;
    }

    .info h3 {
        margin: 0;
    }

    .info img {
        border-radius: 50%;
        height: 65px;
        width: 65px;
    }

    .buttons {
        text-align: right;
    }

    .buttons img {
        height: 35px;
    }

    .single-buttons {
        margin-top: 10px;
        display: none;
        grid-template-columns: 1fr 1fr;
        column-gap: 20px;
        justify-content: center;
        align-items: center;
    }

    .single-buttons a {
        border-radius: 10px;
        display: flex;
        justify-content: center;
        align-items: center;
        color: white;
        font-weight: bold;
        text-decoration: none;
        font-size: 0.9em;
        background: #72D6C9;
        padding: 10px 0;
        transition: filter 0.2s;
    }

    .single-buttons a:hover {
        filter: brightness(0.9);
        transition: filter 0.2s;
    }

    .single-buttons img {
        height: 25px;
        margin-right: 10px;
    }

    @media only screen and (max-device-width: 900px) {
        .buttons {
            display: none;
        }
        .single-buttons {
            display: grid;
        }
    }

    @media only screen and (max-device-width: 640px) {
        .header {
            padding: 15px 20px;
        }
        .single-buttons {
            column-gap: 5px;
        }
    }

</style>
