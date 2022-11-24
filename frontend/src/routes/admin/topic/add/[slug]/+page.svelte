<script lang="ts">

    import {axiosInstance} from "$lib/variables";
    import {goto} from '$app/navigation';
    import {type Topic} from "$lib/model/topic";

    export let data;
    let id: number = parseInt(data.id);

    let topic: Topic = {
        topic: '',
        packID: id
    };

    async function add() {
        const response = await axiosInstance.post('/topic', topic);
        if (response.status === 200) {
            await goto(`/admin/pack/${id}`);
        }
    }
</script>

<svelte:head>
    <title>Thema hinzufügen | Gesprächsthema Generator</title>
</svelte:head>

<div class="container">
    <div class="box">
        <h3>Thema hinzufügen</h3>
        <div class="input">
            <span>Thema</span>
            <textarea bind:value={topic.topic}></textarea>
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

    @media only screen and (max-width: 300px) {
        .box {
            padding: 40px 20px;
        }
    }

</style>