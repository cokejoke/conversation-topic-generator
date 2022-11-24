<script lang="ts">
    import Sidebar from "$lib/components/Sidebar.svelte";
    import PageTransition from "../../lib/components/PageTransition.svelte";
    import {axiosInstance} from "../../lib/variables";
    import {token} from "../../lib/store/token";
    import {goto} from "$app/navigation";
    import {get} from "svelte/store";
    import {onMount} from "svelte";

    /** @type {import('./$types').LayoutData} */
    export let data;

    axiosInstance.interceptors.request.use(
        (config) => {
            if (get(token)) {
                config.headers['Authorization'] = `Bearer ${get(token)}`;
            }
            return config;
        }, (error) => {
            return Promise.reject(error);
        }
    );

    axiosInstance.interceptors.response.use(function (response) {
        return response;
    }, function (error) {
        const code = parseInt(error.response && error.response.status)
        if (code === 401 && data.pathname.split('/')[1] !== 'login') {
            token.set('');
            goto('/admin/login');
        }
        return Promise.reject(error);
    });

    onMount(async () => await axiosInstance.get('/auth'));

</script>

{#if data.pathname.split('/')[2] === 'login'}
    <slot/>
{:else}
    <Sidebar>
        <PageTransition pathname={data.pathname}>
            <div class="content">
                <slot/>
            </div>
        </PageTransition>
    </Sidebar>
{/if}


<style>

    .content {
        padding: 0 40px;
    }

    @media all and (max-device-width: 400px) {
        .content {
            padding: 0 15px;
        }
    }
</style>
