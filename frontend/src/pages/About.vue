<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col d-flex">
        <card title="About" sub-title="Get to know the Constellation Wallet">
          <p>
            Thank you for using the Molly Wallet!
            <br />
            <br />Molly was created by our community member
            <b>Vito</b> and is an application  built on top of infrastructure that was produced by
            the Constellation Engineers. 
            <br />
            <br />The <b>Stardust Collective</b> development team manages all major aspects around the security and development of the wallet, while you, the user, are responsible for securing and backing up your Private Key and password (see
            the
            <b>FAQ</b> for best practices).
            <br />
            <br />

            <b>Something is not working. What should I do?</b>
            <br />If you receive an error message, please visit our <button class="link" @click="openURL('https://t.me/StardustSupport')"><b>Moly Wallet Support</b></button> channel in Telegram.
            <br />
            <br /> Alternatively, you can check to see if this is a known issue on the wallet’s <button class="link" @click="openURL('https://github.com/StardustCollective/molly_wallet/issues')"><b>Github page</b></button>. If it is not a known bug, please create an issue in the provided link with a thorough explanation of the problem. 
            <br /> 
            <br /> Regardless of what error or bug you may encounter, your funds will not be jeopardized by this software.
          </p>
        </card>
      </div>
    </div>
    <div class="row">
      <div class="col d-flex">
        <card title="Frequently Asked Questions">
          <div>
            <b>I thought $DAG was free to transact, why do I need to enter a fee?</b>
            <br />$DAG transactions are indeed free - you can leave this field as 0. The fee is intended for users with higher data throughput needs.
            <br />
            <br />
            <b>How do I securely backup my wallet?</b>
            <br />When creating a new wallet you will be required to create a <b>Private Key (JSON file) </b>and a <b>Password</b>. 
            <br />
            <br />It is important to make a copy of both of these two components, as backup, and keep it in a secure place.
            <br />
            <br /> If you lose your Private Key or your Password, <b>you will not be able to access your wallet</b>.
            <br />
            <br />In summary, the two components below need to be backed up:
            <br />
            <ul>
              <li><b>Private Key (P12 or JSON file)</b></li>
              <li><b>Password</b></li>
            </ul>
          </div>
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
    openURL(url) {
      window.backend.WalletApplication.OpenBrowser(url);
    },
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

<style scoped lang="scss">
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

.link {
  color: #2d9cdb;
  padding: 0px;

  &:hover {
    text-decoration: underline;
  }
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
</style>
