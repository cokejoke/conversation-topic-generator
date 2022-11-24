<script>
    import {goto} from '$app/navigation';
    import {token} from "$lib/store/token";

    function logout() {
        token.set('');
        goto('/admin/login');
    }
</script>

<div class="base">
    <div class="sidebar-wrapper">
        <div class="sidebar">
            <div class="sidebar-top">
                <a class="sidebar-item" href="/admin"><img class="logo" src="/favicon.png" alt="LOGO"></a>
            </div>
            <div class="sidebar-bottom">
                <a class="sidebar-item" href="/admin/pack/add"><img class="icon" src="/svg/add_box.svg" alt="Settings"></a>
                <a class="sidebar-item" on:click={logout}><img class="icon" src="/svg/logout.svg" alt="Logout"></a>
            </div>
        </div>
    </div>
    <div class="content-wrapper">
       <slot/>
    </div>
</div>

<style>

    .logo:hover {
        animation: rotate 2s ease-in-out;
    }

    img {
        height: 45px;
    }

    .icon {
        height: 35px;
    }

    .base {
        display: flex;
        flex-direction: row;
        height: 100vh;
    }

    .sidebar-wrapper {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        align-items: flex-start;
        justify-content: flex-start;
        z-index: 1;
        position: relative;
    }

    .sidebar {
        display: flex;
        flex-direction: column;
        flex-wrap: wrap;
        align-items: center;
        justify-content: space-between;
        height: 100vh;
        width: 70px;
        background: #FFC785;
        box-shadow: rgb(0 0 0 / 8%) 3px 0 5px;
    }

    .sidebar-top, .sidebar-bottom {
        display: flex;
        flex-direction: column;
        flex-wrap: wrap;
        align-items: center;
        justify-content: center;
        width: 100%;
    }

    .sidebar-item {
        cursor: pointer;
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        align-items: center;
        justify-content: center;
        width: 100%;
        height: 80px;
        background-color: transparent;
        transition: background-color .2s ease;
        color: #fff;
        text-decoration: none;
    }

    .sidebar-bottom {
        margin-bottom: 30px;
    }

    .sidebar-bottom .sidebar-item {
        transition: background 0.25s;
    }

    .sidebar-bottom .sidebar-item:hover {
        background: rgba(0, 0, 0, 0.2);
        transition: background 0.25s;
    }

    .content-wrapper {
        display: flex;
        flex: 1 1;
        flex-flow: column;
        overflow-y: auto;
    }

    @media only screen and (max-width: 600px) {
        .base {
            flex-direction: column;
            height: auto;
        }

        .sidebar {
            height: 80px;
            width: 100%;
            display: flex;
            flex-direction: row;
            flex-wrap: nowrap;
            align-items: center;
            justify-content: space-between;
        }

        .sidebar-top, .sidebar-bottom {
            flex-direction: row;
            flex-wrap: nowrap;
            width: auto;
            margin: 0;
        }
        .sidebar-item {
            padding: 0 20px;
            height: 80px;
        }
    }

    @keyframes rotate {
        to {
            transform: rotate(360deg);
        }
    }
</style>