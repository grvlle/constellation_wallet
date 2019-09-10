<template>
    <card title="About" sub-title="Get to know the Constellation Wallet">
        <div>
            <div class="row">
                <div class="col-md-6">
                    <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent odio velit, semper efficitur ultrices sit amet, vestibulum ac ipsum. Sed bibendum tellus non ex tincidunt sagittis. Aenean tristique, est in facilisis malesuada, massa sem
                        dapibus tellus, id cursus quam lacus id nibh. Nulla malesuada orci eget augue ultricies feugiat. Etiam vestibulum rutrum velit, sit amet semper libero luctus a. Morbi tristique urna id purus fermentum, ac egestas leo fringilla. Donec
                        in vestibulum tortor. Suspendisse potenti. Aliquam sit amet sagittis ipsum. Proin cursus id est vitae varius. Donec eleifend, orci nec pellentesque venenatis, tellus sapien tincidunt diam, in cursus elit risus vel tortor. Pellentesque
                        fringilla iaculis elit in viverra. Mauris cursus et ante sit amet hendrerit. Sed sit amet metus scelerisque, aliquet ante non, suscipit turpis. Vestibulum ut ante fringilla est elementum placerat. Nam ac eros eu sapien imperdiet mattis.<br><br>                    Nam interdum diam eu mi fermentum, sit amet porta felis gravida. Etiam molestie nunc vel vestibulum dignissim. Nullam lobortis odio non augue interdum, non consectetur sapien posuere. Duis sed tincidunt lorem. Nulla consectetur risus
                        id sapien pretium, elementum convallis sapien viverra. Suspendisse potenti. Donec sit amet quam tristique orci viverra fringilla vitae vel enim. Nulla laoreet lorem quis gravida pharetra. Ut eu pellentesque magna, sed commodo mi. Fusce
                        sagittis, nisl id mollis egestas, nunc risus euismod sapien, id volutpat sapien purus rhoncus mauris. Proin rutrum interdum erat, sit amet mollis nisl egestas nec. Vestibulum tincidunt pellentesque felis at viverra. Sed quis viverra
                        sem. Nam sodales nulla efficitur finibus sollicitudin. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Praesent vel efficitur lacus. Morbi nec scelerisque magna, mollis porta nibh.<br><br></p>
                </div>
                <div class="col-md-6">
                    <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent odio velit, semper efficitur ultrices sit amet, vestibulum ac ipsum. Sed bibendum tellus non ex tincidunt sagittis. Aenean tristique, est in facilisis malesuada, massa sem
                        dapibus tellus, id cursus quam lacus id nibh. Nulla malesuada orci eget augue ultricies feugiat. Etiam vestibulum rutrum velit, sit amet semper libero luctus a. Morbi tristique urna id purus fermentum, ac egestas leo fringilla. Donec
                        in vestibulum tortor. Suspendisse potenti. Aliquam sit amet sagittis ipsum. Proin cursus id est vitae varius. Donec eleifend, orci nec pellentesque venenatis, tellus sapien tincidunt diam, in cursus elit risus vel tortor. Pellentesque
                        fringilla iaculis elit in viverra. Mauris cursus et ante sit amet hendrerit. Sed sit amet metus scelerisque, aliquet ante non, suscipit turpis. Vestibulum ut ante fringilla est elementum placerat. Nam ac eros eu sapien imperdiet mattis.<br><br>                    Nam interdum diam eu mi fermentum, sit amet porta felis gravida. Etiam molestie nunc vel vestibulum dignissim. Nullam lobortis odio non augue interdum, non consectetur sapien posuere. Duis sed tincidunt lorem. Nulla consectetur risus
                        id sapien pretium, elementum convallis sapien viverra. Suspendisse potenti. Donec sit amet quam tristique orci viverra fringilla vitae vel enim. Nulla laoreet lorem quis gravida pharetra. Ut eu pellentesque magna, sed commodo mi. Fusce
                        sagittis, nisl id mollis egestas, nunc risus euismod sapien, id volutpat sapien purus rhoncus mauris. Proin rutrum interdum erat, sit amet mollis nisl egestas nec. Vestibulum tincidunt pellentesque felis at viverra. Sed quis viverra
                        sem. Nam sodales nulla efficitur finibus sollicitudin. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Praesent vel efficitur lacus. Morbi nec scelerisque magna, mollis porta nibh.<br></p>
                </div>
            </div>
        </div>
    </card>
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

<style>

</style>
