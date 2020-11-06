<template>
  <span
    ><b>Important Update</b><br /><br />

    If you have used the same password for both the KeyStore and the Key
    Password you can now login together with your keyfile (key.p12) and 'a
    single password'.<br /><br />
    If you have used two different passwords you can migrate your Key Alias,
    KeyStore Password and Key Password into 'a single password' by clicking the
    button below.
    <br /><br />
    <!-- <p-button type="danger" style="background: white; border: none;" block @click.native="login()">
      <span style="display: block; color: #B53C19;">MIGRATE</span>
    </p-button> -->
  </span>
</template>

<script>
export default {
  name: "migrate-warning",
};
</script>
