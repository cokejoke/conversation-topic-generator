<script>
    import {createEventDispatcher} from "svelte";

    const dispatch = createEventDispatcher();

    let x, y;
    let ripple = false;
    function onClick(event) {
        dispatch('click');
        if(ripple) {
            return;
        }
        ripple = true;
        x = event.clientX - event.target.offsetLeft;
        y = event.clientY - event.target.offsetTop;
        setTimeout(function () {
            ripple = false;
        }, 500);
    }
</script>

<button on:click={onClick} class="no-select">
    <slot/>
    {#if ripple}
        <span style="top: {y}px; left: {x}px;"></span>
    {/if}
</button>

<style>
    button {
        overflow: hidden;
        display: inline-block;
        position: relative;
        line-height: 1.5em;
        cursor: pointer;
        border: none;
        outline: none;
        background: #DF7599;
        color: rgba(255, 255, 255, 0.87);
        font-weight: bold;
        text-transform: uppercase;
        font-size: 1.2em;
        padding: 20px;
        width: calc(100% - 20px);
        max-width: 500px;
        border-radius: 20px;
        transition: filter 0.25s;
    }

    button:hover {
        filter: brightness(0.95);
        transition: filter 0.25s;
    }

    span {
        position: absolute;
        background: #fff;
        transform: translate(-50%, -50%);
        pointer-events: none;
        border-radius: 50%;
        animation: animate 0.5s linear forwards;
    }

    @keyframes animate {
        0% {
            width: 0;
            height: 0;
            opacity: 0.5;
        }

        100% {
            width: 600px;
            height: 600px;
            opacity: 0;
        }
    }
</style>