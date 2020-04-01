<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col d-flex">
        <card title="About" sub-title="Get to know the Constellation Wallet">
          <p>
            Thank you for using the Constellation Desktop Wallet!
            <br />
            <br />The Constellation Desktop Wallet was built by our community member
            <b>@armedvito</b> and is an application overlay that integrates with the wallet binary produced by
            the Constellation Engineering team. This means that all major aspects around the security of the wallet is managed by the team, as well as yourself - since you will be responsible for handling the wallet backup and private key (see
            the
            <b>FAQ</b> for best practices).
            <br />
            <br />

            <b>Something is not working. What do I do?</b>
            <br />No matter the bug you may run into with the application, keep in mind that
            <b>your funds will never be jeopordized by this software</b>,
            so don't panic if there's an error message.
            <br />
            <br />Please create an issue on GitHub
            <b>github.com/grvlle/constellation_wallet/issues</b> with a thorough explanation of the problem.
            <br />
          </p>
        </card>
      </div>
    </div>
    <div class="row">
      <div class="col d-flex">
        <card title="Frequently Asked Questions">
          <p>
            <b>I thought $DAG was free to transact, why do I need to enter a fee?</b>
            <br />$DAG transactions are indeed free. The fee is intended for companies with higher data throughput demands.
            <br />
            <br />
            <b>How to backup my wallet?</b>
            <br />During wallet creation, you specified the path and key name of your KeyStore File
            <i>(key.p12)</i>. The Private Key will reside in the KeyStore and has three layers of protection. It's up to the user
            in defining the security of this key. The wallet will allow the user to use the same
            <i>Store Password</i> and
            <i>Key Password</i>, but we advise against this for obvious reasons.
            <br />
            <br />One can think of the KeyStore file as a database that hosts multiple private keys. The database is password protected using the (Store Password), without it, no keys inside the
            key śtore can be accessed.
            <br />
            <br />Once the correct KeyStore Password has been provided, there will be a
            <i>Private Key</i> identified by an
            <i>Alias</i> inside. The Private Key will be identified by the
            <i>Alias</i> and decrypted using the
            <i>Private Key Password</i>.
            <br />
            <br />In summary, the four below components needs to be backed up.
            <br />·
            <b>KeyStore File (key.p12)</b>
            <br />·
            <b>Alias</b>
            <br />·
            <b>KeyStore Password</b>
            <br />·
            <b>Key Password</b>
            <br />
            <br />Make sure to store this file somewhere safe, and note down the Store Password, Key Password and Alias on a paper somewhere secure. This means that IF, your storage device containing the Key Store file
            were to be jeopordized by an party with malicous intents, they wouldn't be able to access it without the passwords.
            <!-- Simply navigate to <i>Settings</i> and there you will be presented by two means of backing up your wallet. You can <b>export</b> your keys into an encrypted file on your filesystem. This can later
                        be stored on a USB drive for instance. Make sure that you remove all traces on your computer once en export has been done. <br><br> The other alternative is to copy the <b>Private Key</b> string. Simply click the red button to uncover
                        it, double click to highlight it, then copy and store it somewhere safe. <br><br> The Engineering team is right now investigating means of introducing bip32 encryption which would allow us to generate a more memorable Mnemonic Seed.
            This will then be included in this wallet as well.<br><br>-->

            <br />
          </p>
        </card>
      </div>
    </div>
  </div>
</template>

<script>
import NotificationTemplate from "./Notifications/WalletCopied";

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
$color-github: #f05033;
$color-twitter: #2b97f1;
$color-dribbble: #ef5a92;
$color-instagram: #527fa6;
$color-linkedin: #006599;
$color-facebook: #3b5a9b;
$color-telegram: #0088cc;
$color-discord: #7289da;
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
  width: 5em;
  height: 5em;
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
    box-shadow: 0em 0.0625em 0.1875em rgba(0, 0, 0, 0.5) inset;
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
    top: 0.0625em;
  }
}

/* Tooltips */

.tooltip {
  display: block;
  position: absolute;
  top: 0;
  left: 50%;
  padding: 0.8rem 1rem;
  border-radius: 2.5em;
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
    bottom: 0.0625em;
    left: 50%;
    width: 0;
    height: 0;
    content: "";
    border: solid;
    border-width: 0.625em 0.625em 0 0.625em;
    border-color: transparent;
    transform: translate(-50%, 100%);
  }
}

.card {
  width: 100%;
  display: flex;
  flex-direction: column;
}

.card-footer {
  margin-top: auto;
}
</style>
