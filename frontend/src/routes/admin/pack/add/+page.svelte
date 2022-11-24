<script lang="ts">

    import {type Pack} from "../../../../lib/model/pack";
    import {axiosInstance} from "../../../../lib/variables";
    import {goto} from '$app/navigation';

    let pack: Pack = {
        name: '',
        author: '',
        language: 'DE'
    };

    async function add() {
        const response = await axiosInstance.post('/pack', pack);
        if (response.status === 200) {
            await goto('/admin');
        }
    }
</script>

<svelte:head>
    <title>Pack hinzufügen | Gesprächsthema Generator</title>
</svelte:head>

<div class="container">
    <div class="box">
        <h3>Pack hinzufügen</h3>
        <div class="row">
            <div class="input">
                <span>Name</span>
                <input bind:value={pack.name} type="text" required>
            </div>
            <div class="input">
                <span>Autor</span>
                <input bind:value={pack.author} type="text" required>
            </div>
        </div>
        <div class="input">
            <span>Bild URL (optional)</span>
            <input bind:value={pack.imageUrl} type="text">
        </div>
        <button on:click={add}>Speichern</button>
    </div>
</div>

<style>

    h3 {
        text-align: center;
        color: rgba(0, 0, 0, 0.6);
        margin: 0 0 40px;
        font-size: max(1.2em, 18px);
    }

    .container {
        display: flex;
        width: 100%;
        height: 100vh;
        align-items: center;
        justify-content: center;
    }

    .box {
        background: rgba(255, 255, 255, 0.8);
        padding: 40px;
        border-radius: 16px;
        width: 100%;
        max-width: 500px;
    }

    .row {
        display: grid;
        grid-template-columns: 1fr 1fr;
        column-gap: 35px;
    }

    .input {
        margin-bottom: 35px;
    }

    .input span {
        display: block;
        color: #333333;
        font-size: 14px;
        margin-left: 10px;
        margin-bottom: 5px;
    }

    .input input {
        outline: none;
        border-color: rgba(0, 0, 0, 0.15);
        border-radius: 7px;
        display: block;
        width: calc(100% - 40px);
        padding: 10px 20px;
    }

    button {
        background: #DF7599;
        border: none;
        cursor: pointer;
        width: 100%;
        padding: 15px 0;
        color: white;
        font-size: 18px;
        font-weight: bold;
        text-transform: uppercase;
        border-radius: 12px;
        transition: opacity 0.25s;
    }

    button:hover {
        opacity: 0.8;
        transition: opacity 0.25s;
    }

    @media only screen and (max-width: 600px) {
        .row {
            grid-template-columns: auto;
        }
    }

    @media only screen and (max-width: 300px) {
        .box {
            padding: 40px 20px;
        }
    }

</style>