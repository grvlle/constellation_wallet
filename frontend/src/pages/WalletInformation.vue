<template>
  <div id="app" class="row">
    <div class="col-xl-6 col-lg-7 col-md-6">
      <user-card></user-card>
    </div>
    <div class="col-xl-6 col-lg-7 col-md-6">
      <edit-profile-form></edit-profile-form>
    </div>
  </div>
</template>
<script>
import EditProfileForm from "./UserProfile/EditProfileForm.vue";
import UserCard from "./UserProfile/UserCard.vue";
export default {
  components: {
    EditProfileForm,
    UserCard
  }
};
</script>
<style>
</style>
