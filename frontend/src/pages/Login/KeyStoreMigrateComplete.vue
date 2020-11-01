<template>
  <div class="container">
    <div class="row">
      <div class="col-12">
        <form ref="textareaform" @submit.prevent="form" class="container">
          <div class="row">
            <div class="col mx-auto login-box">
              <div class="input-box">
                <p>A new file has been created and ready to be used for Molly Wallet 2.0. This file is located in the same directory as the original.</p>
                <p style="font-weight: bold">{{ filePath }}.</p>
              </div>
              <br />
              <div class="input-box">
                <svg
                  class="checkmark"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 52 52"
                >
                  <circle
                    class="checkmark__circle"
                    cx="26"
                    cy="26"
                    r="25"
                    fill="none"
                  />
                  <path
                    class="checkmark__check"
                    fill="none"
                    d="M14.1 27.2l7.1 7.2 16.7-16.8"
                  />
                </svg>
              </div>
              <div class="button-box">
                <div class="container">
                  <div class="row">
                    <div class="col">
                      <p-button
                        class="btn-secondary"
                        block
                        @click.native="moveToLogin()"
                      >
                        <span style="display: block"> NEXT</span>
                      </p-button>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </form>
      </div>
    </div>
    <page-overlay text="Loading..." :isActive="overlay" />
  </div>
</template>

<script>
import { mapState } from "vuex";

export default {
  components: {},
  name: "keystore-migration-complete",
  data: () => ({
    keystorePassword: "",
    KeyPassword: "",
    overlay: false,
  }),
  computed: {
    ...mapState("app", ["network"]),
    keystorePath: {
      get() {
        return this.$store.state.wallet.keystorePath;
      },
      set(value) {
        this.$store.commit("wallet/setKeystorePath", value);
      },
    },
    alias: {
      get() {
        return this.$store.state.wallet.alias;
      },
      set(value) {
        this.$store.commit("wallet/setAlias", value);
      },
    },
    filePath: {
      get() {
        return this.$route.params.filePath;
      },
    },
  },
  methods: {
    moveToLogin: function() {
      this.$router.push({
        name: "login single password",
        params: {
          message:
            "Please enter the credentials to your Private Key file.",
        },
      });
    },
  },
};
</script>

<style scoped lang="scss">
.login-box {
  max-width: 29rem;
  min-width: 29rem;
  padding-bottom: 2rem;
}

.btn-secondary {
  background: #db6e44;
  border: 0.0625rem solid #db6e44;
  font-family: Poppins;
  font-style: normal;
  font-weight: 500;
  font-size: 0.75rem;
  line-height: 1.125rem;
  text-align: center;
  letter-spacing: 0.1em;
  border-radius: 0.25rem;
  color: #ffffff;

  &:hover {
    background: #af5836;
    border: 0.0625rem solid #af5836;
  }

  &:active,
  &:focus {
    background: #db6e44 !important;
    border: 0.0625rem solid #db6e44 !important;
  }

  &:active {
    outline-color: #db6e44 !important;
    outline-width: 0;
  }

  &:disabled {
    background: #e9a88f !important;
    border: 0.0625rem solid #e9a88f !important;
  }
}

div.input-box {
  font-family: Poppins;
  font-style: normal;
  font-weight: 500;
  font-size: 10.25rem;
  line-height: 20.25rem;
  color: #666666;
}

.button-box .container {
  margin-left: 0em;
  margin-right: 0em;
  padding-left: 0em;
  padding-right: 0em;
}

.button-box .container .row {
  margin-left: 0em;
  margin-right: 0em;
  padding-left: 0em;
  padding-right: 0em;
  margin-top: 1.25em;
}

.button-box .container .row [class^="col"] {
  margin-left: 0em;
  margin-right: 0em;
  padding-left: 0em;
  padding-right: 0em;
}

.checkmark__circle {
  stroke-dasharray: 166;
  stroke-dashoffset: 166;
  stroke-width: 2;
  stroke-miterlimit: 10;
  stroke: #db6e44;
  fill: none;
  animation: stroke 0.6s cubic-bezier(0.65, 0, 0.45, 1) forwards;
}

.checkmark {
  width: 11rem;
  height: 11rem;
  border-radius: 50%;
  display: block;
  stroke-width: 2;
  stroke: #fff;
  stroke-miterlimit: 10;
  margin: 10% auto 30%;
  box-shadow: inset 0 0 0 #db6e44;
  animation: fill 0.4s ease-in-out 0.4s forwards,
    scale 0.3s ease-in-out 0.9s both;
}

.checkmark__check {
  transform-origin: 50% 50%;
  stroke-dasharray: 48;
  animation: stroke 0.3s cubic-bezier(0.65, 0, 0.45, 1) 0.8s forwards;
}

@keyframes stroke {
  100% {
    stroke-dashoffset: 0;
  }
}
@keyframes scale {
  0%,
  100% {
    transform: none;
  }
  50% {
    transform: scale3d(1.1, 1.1, 1);
  }
}
@keyframes fill {
  100% {
    box-shadow: inset 0 0 0 90 #db6e44;
  }
}
</style>
