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

app.mount('#app');

