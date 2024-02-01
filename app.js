import * as bootstrap from 'bootstrap';
window.bootstrap = bootstrap;
import 'bootstrap';
import "/css/styles.scss";


import { createApp } from 'vue';

const app = createApp({});

import HomeComponent from "./vueComponents/Home.vue";
app.component('home', HomeComponent);
import ConfirmComponent from "./vueComponents/Confirm.vue";
app.component('confirm', ConfirmComponent);
import LeftComponent from "./vueComponents/LeftMenu.vue";
app.component('left-menu', LeftComponent);
import TopComponent from "./vueComponents/Top.vue";
app.component('top', TopComponent);
import ShopComponent from "./vueComponents/Shop.vue";
app.component('shop', ShopComponent);
import MapsComponent from "./vueComponents/Maps.vue";
app.component('maps', MapsComponent);
import WarehouseComponent from "./vueComponents/Warehouse.vue";
app.component('warehouse', WarehouseComponent);
import UsersComponent from "./vueComponents/Users.vue";
app.component('users', UsersComponent);

app.mount('#app');

