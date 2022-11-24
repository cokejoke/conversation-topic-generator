<script lang="ts">
    import { fly, fade } from 'svelte/transition';
    export let pathname: string = '';
</script>

{#key pathname}
    <div style="height: 100%; width: 100%;" in:fly={{ x: -15, duration: 500, delay: 500 }} out:fade={{ duration: 200 }}>
        <slot />
    </div>
{/key}