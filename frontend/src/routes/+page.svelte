<script lang="ts">
    import {LottiePlayer} from '@lottiefiles/svelte-lottie-player';
    import {fade} from 'svelte/transition';
    import {onMount} from "svelte";
    import Bubble from "../lib/components/Bubble.svelte";
    import axios from "axios";
    import Button from "../lib/components/Button.svelte";
    import { env } from '$env/dynamic/public';

    let loading = false;
    let displayTopic = false;

    let seen = [];

    let topic;

    async function loadTopic(initial = false) {
        if (!initial && (loading || !displayTopic)) {
            return;
        }
        if (!displayTopic) loading = true;
        displayTopic = false;

        try {
            let response = await axios.get(`${'http://127.0.0.1:8991'}/topic/random`);
            while (seen.includes(response.data.id)) {
                response = await axios.get(`${'http://127.0.0.1:8991'}/topic/random`);
            }
            seen.push(response.data.id);
            topic = response.data.topic;
            setTimeout(() => loading = false, 1000);
        } catch (error) {
            console.log(error);
        }
    }


    const spaceCode = 32;

    let holdingSpace = false;

    function handleKeydown(event) {
        if (event.keyCode === spaceCode) {
            if (holdingSpace || !displayTopic) {
                return;
            }
            holdingSpace = true;
            loadTopic();
        }
    }

    function handleKeyup(event) {
        if (event.keyCode === spaceCode) {
            holdingSpace = false;
        }
    }

    onMount(async () => loadTopic(true));


</script>

<svelte:head>
    <title>Gesprächsthema Generator</title>
</svelte:head>

<svelte:window on:keydown={handleKeydown} on:keyup={handleKeyup}/>
<main>
        <h1 class="no-select">Gesprächsthema<br>Generator</h1>
        <div class="content-wrapper">
            <div>
                {#if loading}
                    <div in:fade out:fade={{duration: 1500}} on:outroend="{() => displayTopic = true}">
                        <LottiePlayer
                                src="/catloader.json"
                                autoplay="{true}"
                                loop="{true}"
                                controls="{false}"
                                controlsLayout="{[]}"
                                renderer="svg"
                                background="transparent"
                                height="{300}"
                                width="{300}"
                        />
                    </div>
                {/if}
                {#if displayTopic}
                    <div in:fade out:fade on:outroend="{() => loading = true}">
                        <Bubble>{topic}</Bubble>
                    </div>
                {/if}
            </div>
        </div>
        <Button on:click={() => loadTopic()}>Neues Thema losen<br>[LEERTASTE]</Button>
</main>

<style>
    main {
        height: calc(100% - 40px);
        text-align: center;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        padding: 20px 0;
    }

    .wrapper {
        padding: 40px 0;
    }

    h1 {
        text-align: center;
        text-transform: uppercase;
        margin: 0;
    }

    .content-wrapper {
        display: flex;
        justify-content: center;
        align-items: center;
        max-height: 500px;
        height: 100%;
    }

</style>
