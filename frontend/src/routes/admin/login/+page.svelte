<script lang="ts">
    import {goto} from '$app/navigation';
    import {onMount} from 'svelte';
    import {scale} from 'svelte/transition';
    import {token} from "$lib/store/token";
    import {get} from "svelte/store";
    import {axiosInstance} from "$lib/variables";

    let show = false;
    onMount(() => {
        if (get(token) !== "") {
            goto("/admin");
        }
        show = true;
    });

    let loginRequest = {
        username: '',
        password: '',
    }

    async function login() {
        try {
            let response = await axiosInstance.post('/login', loginRequest);
            if (response.status == 200 && response.data) {
                token.set(response.data.token);
                await goto("/admin");
            }
        } catch (error) {
            if (error.response.status == 404) {
                displaySnackbar("Falscher Benutzername!");
            } else {
                displaySnackbar("Falsches Passwort!");
            }
        }
    }

    let showSnackbar: boolean = false;
    let snackbarMsg = "Test";

    function displaySnackbar(msg: string): void {
        snackbarMsg = msg;
        showSnackbar = true;
        setTimeout(() => showSnackbar = false, 3000);
    }

</script>

<svelte:head>
    <title>Login | Gesprächsthema Generator</title>
</svelte:head>

<div class="wrapper">
    {#if show}
        <form on:submit|preventDefault={login} in:scale={{ duration: 1000 }} id="login-box">
            <input bind:value={loginRequest.username} class="login-input" type="text" id="username" name="username"
                   placeholder="Benutzername"><br>
            <input bind:value={loginRequest.password} class="login-input" type="password" id="password" name="password"
                   placeholder="Passwort"><br>
            <div>
                <button type="submit" id="login-button">Anmelden</button>
            </div>
        </form>
    {/if}

    <div id="snackbar" class:show={showSnackbar}>{snackbarMsg}</div>
</div>

<style>

    .wrapper {
        display: flex;
        width: 100%;
        height: 100vh;
        align-items: center;
        justify-content: center;
    }

    #login-box {
        display: inline-block;
        background: #FFC785;
        border-radius: 20px;
        width: calc(100% - 100px);
        max-width: 400px;
        height: 200px;
        padding: 50px;
        margin-bottom: 100px;
    }

    .login-input {
        outline: none;
        display: inline-block;
        width: calc(100% - 80px);
        border: none;
        border-radius: 17px;
        background: white;
        font-size: 17px;
        font-weight: bold;
        padding: 15px 40px;
        margin-bottom: 20px;
        color: rgba(0, 0, 0, 0.8);
    }

    .login-input::placeholder {
        color: rgba(0, 0, 0, 0.5);
    }

    #login-footer div:last-child {
        text-align: right;
    }

    #login-button {
        cursor: pointer;
        border: none;
        border-radius: 17px;
        font-size: 15px;
        font-weight: bold;
        padding: 12px 25px;
        color: white;
        background: #72D6C9;
        transition: background-color 0.5s;
        width: 100%;
        text-transform: uppercase;
    }

    #login-button:hover {
        transition: background-color 0.5s;
        background: #72D6C9;
    }

    /* Snackbar */

    #snackbar {
        visibility: hidden;
        min-width: 250px;
        margin-left: -125px;
        background-color: #333;
        color: #fff;
        text-align: center;
        border-radius: 2px;
        padding: 16px;
        position: fixed;
        z-index: 1;
        left: 50%;
        bottom: 30px;
        font-family: sans-serif;
    }

    #snackbar.show {
        visibility: visible;
        animation: fadein 0.5s, fadeout 0.5s 2.5s;
    }

    @keyframes fadein {
        from {
            bottom: 0;
            opacity: 0;
        }
        to {
            bottom: 30px;
            opacity: 1;
        }
    }

    @keyframes fadeout {
        from {
            bottom: 30px;
            opacity: 1;
        }
        to {
            bottom: 0;
            opacity: 0;
        }
    }

    @media only screen and (max-width: 400px) {
        #login-box {
            padding: 30px;
        }
    }
</style>