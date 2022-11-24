<script lang="ts">

    import {type Topic} from "$lib/model/topic";
    import { fade } from 'svelte/transition';

    export let topic: Topic;

</script>

<a href="/admin/topic/edit/{topic.id}" in:fade class="topic">
    <span>{topic.topic}</span>
</a>

<style>
    .topic {
        text-decoration: none;
        display: block;
        cursor: pointer;
        background: white;
        padding: 15px 30px;
        margin-top: 20px;
        border-radius: 10px;
        transition: background 0.5s;
    }

    .topic:hover {
        transition: background 0.5s;
        background: rgba(255, 255, 255, 0.8);
    }

    span {
        color: #333333;
    }
</style>