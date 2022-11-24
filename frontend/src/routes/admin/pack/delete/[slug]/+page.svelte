<script lang="ts">

    import {axiosInstance} from "$lib/variables";
    import {goto} from '$app/navigation';

    export let data;
    let id: number = parseInt(data.id);

    async function remove() {
        const response = await axiosInstance.delete(`/pack/${id}`);
        if (response.status === 200) {
            await goto(`/admin`);
        }
    }
</script>

<svelte:head>
    <title>Thema bearbeiten | Gesprächsthema Generator</title>
</svelte:head>

<div class="container">
    <div class="box">
        <h3>Pack löschen</h3>
        <div class="input">
            <p>Bist du dir sicher das du das Pack löschen willst?</p>
        </div>
        <div class="row">
            <button on:click={async () => await goto(`/admin/pack/${id}`)}>Abbrechen</button>
            <button on:click={remove}>Löschen</button>
        </div>
    </div>
</div>

<style>

    p {
        text-align: center;
        color: #333333;
    }

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

    .input textarea {
        outline: none;
        border-color: rgba(0, 0, 0, 0.15);
        border-radius: 7px;
        display: block;
        width: calc(100% - 40px);
        height: 60px;
        padding: 10px 20px;
    }

    .row {
        display: grid;
        grid-template-columns: 1fr 1fr;
        column-gap: 20px;
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

    @media only screen and (max-width: 500px) {
        .row {
            grid-template-columns: auto;
        }

        button:first-child {
            margin-bottom: 20px;
        }
    }

    @media only screen and (max-width: 300px) {
        .box {
            padding: 40px 20px;
        }
    }

</style>