<template>
    <div id="app">
        <div class="row">
            <div class="col-md-6">
                <card title="About" sub-title="Get to know the Constellation Wallet">
                    <p>Thank you for using the Constellation Desktop Wallet!<br><br> The Constellation Desktop Wallet was built by our community member <b>@armedvito</b> and is an application overlay that integrates with the wallet binary produced by
                        the Constellation Engineering team. This means that all major aspects around the security of the wallet is managed by the team, as well as yourself - since you will be responsible for handling the wallet backup and private key (see
                        the <b>FAQ</b> for best practices). <br><br> This is the first iterration of the Desktop Wallet, with many more versions and alternatives to come. If you were to experience any errors with it (you will), or have any feature
                        requests - then please reach out to <b>@armedvito</b> on <b>Telegram</b> or <b>Twitter</b> DMs. No matter the bug you may run into with the application, keep in mind that <b>your funds will never be jeopordized by this software</b>,
                        so don't panic if there's an error message.<br><br>
    
                    </p>
                </card>
    
                <card>
                    <div class="social-icons" style="align-items: center; justify-content: center;">
    
                        <a class="social-icon social-icon--github">
                                    <i class="fab fa-github" style="color: #fff;"></i>
                                    <div class="tooltip">Github</div>
                                  </a>
                        <a class="social-icon social-icon--twitter">
                                    <i class="fab fa-twitter" style="color: #fff;"></i>
                                    <div class="tooltip">Twitter</div>
                                  </a>
                        <a class="social-icon social-icon--instagram" style="vertical-align: middle; background: #d6249f; background: radial-gradient(circle at 30% 107%, #fdf497 0%, #fdf497 5%, #fd5949 45%,#d6249f 60%,#285AEB 90%);">
                                    <i class="fab fa-instagram" style="color: #fff;"></i>
                                    <div class="tooltip">Instagram</div>
                                  </a>
                        <a class="social-icon social-icon--telegram">
                                    <i class="fab fa-telegram-plane" style="color: #fff;"></i>
                                    <div class="tooltip">Telegram</div>
                                  </a>
                        <a class="social-icon social-icon--discord">
                                    <i class="fab fa-discord" style="color: #fff;"></i>
                                    <div class="tooltip">Discord</div>
                                  </a>
                        <a class="social-icon social-icon--website">
                                    <i class="fas fa-globe-americas" style="color: #fff;"></i>
                                    <div class="tooltip">Website</div>
                                  </a>
    
                        <br>
                    </div>
                </card>
            </div>
            <div class="col-md-6">
                <card title="Frequently Asked Questions">
    
                    <p>
                        <b>How to backup my wallet?</b><br> Simply navigate to <i>Settings</i> and there you will be presented by two means of backing up your wallet. You can <b>export</b> your keys into an encrypted file on your filesystem. This can later
                        be stored on a USB drive for instance. Make sure that you remove all traces on your computer once en export has been done. <br><br> The other alternative is to copy the <b>Private Key</b> string. Simply click the red button to uncover
                        it, double click to highlight it, then copy and store it somewhere safe. <br><br> The Engineering team is right now investigating means of introducing bip32 encryption which would allow us to generate a more memorable Mnemonic Seed.
                        This will then be included in this wallet as well.<br><br>
    
                        <b>Something is not working. What do I do?</b><br> Please send an e-mail to <b>@armedvito</b> with a thorough explanation of the problem. Also, please include the <b>wallet.log</b> file located in your
                        <i>.dag</i> folder.
                        <br><br>
                        <b>I thought $DAG was free to transact, why do I need to enter a fee?</b><br> $DAG transactions are indeed free. The fee is intended for companies with higher data throughput demands.
                    </p>
    
                </card>
            </div>
        </div>
    </div>
</template>

<script>
import NotificationTemplate from './Notifications/WalletCopied';

export default {
    data() {
        return {
            type: ["", "info", "success", "warning", "danger"],
            notifications: {
                topCenter: false
            }
        };
    },
    methods: {
        notifyVue(verticalAlign, horizontalAlign) {
            const color = Math.floor(Math.random() * 4 + 1);
            this.$notify({
                component: NotificationTemplate,
                icon: "ti-gift",
                horizontalAlign: horizontalAlign,
                verticalAlign: verticalAlign,
                type: this.type[color]
            });
        }
    }
};
</script>

<style lang="scss">
/* Color Variables */

$color-codepen: #000;
$color-website: #000;
$color-github: #F05033;
$color-twitter: #2b97f1;
$color-dribbble: #ef5a92;
$color-instagram: #527fa6;
$color-linkedin: #006599;
$color-facebook: #3b5a9b;
$color-telegram: #0088CC;
$color-discord: #7289DA;
/* Social Icon Mixin */

@mixin social-icon($color) {
    background: $color;
    color: #fff;
    .tooltip {
        background: $color;
        color: #fff;
        &:after {
            border-top-color: $color;
        }
    }
}

/* Social Icons */

.social-icons {
    display: flex;
}

.social-icon {
    display: flex;
    align-items: center;
    justify-content: center;
    position: relative;
    width: 80px;
    height: 80px;
    margin: 0 0.5rem;
    border-radius: 50%;
    cursor: pointer;
    font-family: "Helvetica Neue", "Helvetica", "Arial", sans-serif;
    font-size: 2.5rem;
    text-decoration: none;
    transition: all 0.15s ease;
    &:hover {
        color: #fff;
        .tooltip {
            visibility: visible;
            opacity: 1;
            transform: translate(-50%, -150%);
        }
    }
    &:active {
        box-shadow: 0px 1px 3px rgba(0, 0, 0, 0.5) inset;
    }
    &--linkedin {
        @include social-icon($color-linkedin);
    }
    &--website {
        @include social-icon($color-website);
    }
    &--twitter {
        @include social-icon($color-twitter);
    }
    &--codepen {
        @include social-icon($color-codepen);
    }
    &--facebook {
        @include social-icon($color-facebook);
    }
    &--discord {
        @include social-icon($color-discord);
    }
    &--telegram {
        @include social-icon($color-telegram);
    }
    &--instagram {
        @include social-icon($color-instagram);
    }
    &--dribbble {
        @include social-icon($color-dribbble);
    }
    &--github {
        @include social-icon($color-github);
    }
    i {
        position: relative;
        top: 1px;
    }
}

/* Tooltips */

.tooltip {
    display: block;
    position: absolute;
    top: 0;
    left: 50%;
    padding: 0.8rem 1rem;
    border-radius: 40px;
    font-size: 0.8rem;
    font-weight: bold;
    opacity: 0;
    pointer-events: none;
    text-transform: uppercase;
    transform: translate(-50%, -100%);
    transition: all 0.3s ease;
    z-index: 1;
    &:after {
        display: block;
        position: absolute;
        bottom: 1px;
        left: 50%;
        width: 0;
        height: 0;
        content: "";
        border: solid;
        border-width: 10px 10px 0 10px;
        border-color: transparent;
        transform: translate(-50%, 100%);
    }
}
</style>
