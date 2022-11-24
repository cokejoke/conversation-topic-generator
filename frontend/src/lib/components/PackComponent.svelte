<script lang="ts">

    import {type Pack} from "$lib/model/pack";
    import { fade } from 'svelte/transition';

    export let pack: Pack;

</script>

<a in:fade href="/admin/pack/{pack.id}" class="pack">
    <div class="row">
        <img src="{pack.imageUrl ? pack.imageUrl : '/placeholder.png'}" alt="Cover">
        <div>
            <h4 class="title">{pack.name}</h4>
            <span class="author">{pack.author}</span>
        </div>
    </div>
</a>

<style>
    .pack {
        text-decoration: none;
        background: rgba(255, 255, 255, 0.85);
        border-radius: 30px;
        margin-bottom: 40px;
        transition: opacity 0.5s;
        padding: 10px;
    }

    .row {
        display: grid;
        grid-template-columns: 75px auto;
        column-gap: 15px;
        align-items: center;
    }

    .pack:hover {
        opacity: 0.8;
        transition: opacity 0.5s;
    }

    img {
        border-radius: 50%;
        height: 75px;
        width: 75px;
    }

    .title {
        color: #333333;
        margin: 0;
        font-size: max(0.9em, 13px);
    }

    .author {
        font-size: max(0.7em, 11px);
        color: #4f4f4f;
    }


</style>