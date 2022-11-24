<div class="thought">
    <slot/>
</div>

<style>
    .thought {
        display: flex;
        background-color: #fff;
        padding: 50px;
        border-radius: 30px;
        width: calc(100% - 140px);
        max-width: 666px;
        min-height: 40px;
        margin: 20px;
        position: relative;
        align-items: center;
        justify-content: center;
        text-align: center;
        color: rgba(0, 0, 0, 0.7);
        font-size: 1.5em;
    }

    .thought:before,
    .thought:after {
        content: "";
        background-color: #fff;
        border-radius: 50%;
        display: block;
        position: absolute;
        z-index: -1;
    }

    .thought:before {
        width: 44px;
        height: 44px;
        top: -12px;
        left: 28px;
        box-shadow: -50px 30px 0 -12px #fff;
    }

    .thought:after {
        bottom: -10px;
        right: 26px;
        width: 30px;
        height: 30px;
        box-shadow: 40px -34px 0 0 #fff,
        -28px -6px 0 -2px #fff,
        -24px 17px 0 -6px #fff,
        -5px 25px 0 -10px #fff;
    }

    @media all and (max-device-width: 640px) {
        .thought {
            padding: 40px;
        }
    }
</style>