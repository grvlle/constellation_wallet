<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col-md-6 d-flex">
        <user-card></user-card>
      </div>
      <div class="col-md-6 d-flex">
        <edit-profile-form></edit-profile-form>
      </div>
    </div>
  </div>
</template>
<script>
import EditProfileForm from "./UserProfile/EditProfileForm.vue";
import UserCard from "./UserProfile/UserCard.vue";
export default {
  components: {
    EditProfileForm,
    UserCard,
  },
};
</script>

<style scoped lang="scss">
.container {
  padding: 0em;
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
