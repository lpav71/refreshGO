<template>
    <div class="main">
        <div class="left">
            <div style="" class="user-select">
                <span style="margin-bottom: 20px">Выберите пользователя</span>
                <div style="margin-top: -33px;">
                    <i class="fa fa-search lupa2"></i>
                    <input class="select_user" type="text" @keydown="searchClientToKey" @input="searchClients" v-model="searchQuery" placeholder="Поиск логина" />
                </div>
                <div class="popup" v-show="findWindow" ref="findBlock">
                    <ul>
                        <li v-for="(f, i) in filteredClients" @click="selectFindClient(i)">{{ f.login }}</li>
                    </ul>
                </div>
                <div class="user_name">
                    <img src="images/shop/ava.svg" style="width: 22px;margin-top: 10px;margin-left: 10px;"  alt=""/>
                    <i class="fa fa-star" style="margin-left: 119px;position: relative;top: 5px;left: 71px;"></i>
                    <div style="display: inline-block;position: relative;top: -30px;left: 28px;">
                        <span class="client_login">{{ findClient.Login }}</span>
                    </div>
                </div>
                <div class="info" v-if="findClientStatus">
                    <img src="images/shop/rub.png" width="16" height="16" style="margin-top: 12px;margin-left: 17px;"  alt=""/>
                    <span class="rub_summa">{{ findClient.Amount }}</span>
                    <img src="images/shop/coin.png" width="16" height="16" style="margin-top: 10px;margin-left: 97px;"  alt=""/>
                    <span class="rub_summa" style="color: #FFCD00;">{{ findClient.Bonus }}</span>
                </div>
            </div>
            <button class="btn btn-success" @click="clearClient" type="button" style="width: 250px;">Нет логина [гость]</button>
            <span style="margin-bottom: -24px;margin-top: 8px;">Корзина</span>
            <div class="cart">
                <div v-for="(c, index) in cart" class="cart_element" :class="{'goods-discount': c.allowDiscount}">
                    <span style="padding-right: 0; opacity: 0.2">x{{ c.qty }}</span>
                    <div class="cart_name">
                        <span>{{ c.name }}</span>
                    </div>
                    <i class="fa fa-trash" style="cursor: pointer; opacity: 0.2" @click="del(index)"></i>
                </div>
            </div>
            <div class="bottom3"><span>Применить промокод</span>
                <input v-model="promo_code" class="user_name" type="text" placeholder="refresh 2023" style="text-align: center;" />
<!--                <div style="display: flex;justify-content: space-between;width: 76%;align-items: center;">-->
<!--                    <a href="" @click="verify_promo_code">Проверить промокод</a><i v-if="verifyPromo" class="fa fa-thumbs-up promo-ico"></i><i v-if="!verifyPromo" class="fa fa-times promo-ico"></i>-->
<!--                </div>-->
                <span>Итого</span>
                <div class="user_name">
                    <img src="images/shop/rub.png" width="16" height="16" style="margin-top: 12px;margin-left: 17px;"  alt=""/>
                    <span class="rub_summa">{{ rubs }}</span>
                    <img src="images/shop/coin.png" width="16" height="16" style="margin-top: 10px;margin-left: 97px;"  alt=""/>
                    <span class="rub_summa" style="color: #FFCD00;">{{ bonuses }}</span>
                </div>
            </div>
            <button class="btn btn-danger mt-4" type="button" style="width: 250px;" @click="pay">Оплатить</button>
        </div>
        <div class="right">
            <div class="top"><img src="images/shop/2.svg"  alt=""/><span style="margin-left: 14px;">Магазин</span><i class="fa fa-search lupa"></i><input @keydown="enter" v-model="product_name" class="input_shop" type="text" placeholder="Поиск товаров" />
                <button class="btn btn-success butt" type="button">Все</button>
                <button class="btn btn-success butt" type="button">Товары</button>
                <button class="btn btn-success butt" type="button">Пакетные предложения</button>
                <button class="btn btn-success butt" type="button">Услуги</button>
            </div>
            <div class="fon">
                <div class="bottom2" style="border-radius: 20px;padding: 30px;">
                    <div class="goods" v-for="(g, index) in goods">
                        <div class="goods_img"><img :src="g.Icon" height="170"  alt=""/></div>
                        <div class="info-block">
                            <div class="goods_text">
                                <div>
                                    <div style="opacity: 0.5">{{ g.Name }}</div>
                                    <div style="font-size: 12px;height: 17px;overflow: hidden;">{{ g.Product }} {{ g.ProductParam }}</div>
                                </div>
                                <div v-if="g.Price != 0" style="margin-top: 25px"><span>{{ g.Price }} Р</span></div>
                            </div>
                            <div class="button-block">
                              <button class="goods_cart" @click="toCartPlus(index)" type="button" style="background: linear-gradient(180deg, rgba(40, 79, 98, 0.05) 0%, #183F52 100%);margin-bottom: 3px;"><div class="pm" style="position: relative;right: 2.4px;"><span>+</span></div></button>
                                <button class="goods_cart" @click="toCartMinus(index)" type="button"><div class="pm" style="position: relative;bottom: 4px;"><span>-</span></div></button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <!-- Модальное окно -->
    <div class="modal fade" id="payModal" tabindex="-1" aria-labelledby="payModalLabel" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered">
            <div class="modal-content content">
                <div class="modal-header">
                    <h5 class="modal-title" id="payModalLabel">Оплата</h5>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Закрыть"></button>
                </div>
                <div class="modal-body">
                    <div class="main2">
                        <div class="div_2">
                            <div class="div_data">
                                <div class="div_3">
                                    <div>
                                        <span>Способы оплаты</span>
                                        <div class="div_4"><input id="formcheck_1" class="formcheck_1" name="pay" type="radio" checked /><label class="form-label" for="formcheck_1">Наличные</label></div>
                                        <div class="div_4"><input id="formcheck_2" class="formcheck_1" name="pay" type="radio" /><label class="form-label" for="formcheck_2">Карта</label></div>
                                    </div>
                                </div>
                                <div class="div_3">
                                    <div>
                                        <span>Чек</span>
                                        <div class="div_4"><input id="formcheck_3" class="formcheck_1" type="checkbox" /><label class="form-label" for="formcheck_3">Отправить на почту</label></div>
                                        <div class="div_4"><input id="formcheck_4" class="formcheck_1" type="checkbox" /><label class="form-label" for="formcheck_4">Печать</label></div>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div class="div_7" style="border-style: none;">
                            <div class="bottom"><span class="span_1">Комментарий к платежу</span><input class="div_6" type="text" /></div>
                        </div>
                    </div>
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn btn-primary save" @click="payJson">Оплата</button>
                </div>
            </div>
        </div>
    </div>

<!--    <message ref="message"></message>-->
<!--    <system-message ref="sysMessage"></system-message>-->

</template>

<script>

export default {
    name: "shop",
    props: ['club_id', 'user_id'],
    data() {
        return {
            goods: [],
            product_name: '',
            cart: [],
            cartElement: {},
            rubs: 0,
            bonuses: 0,
            findClient: {},
            findClientStatus: false,
            modal: null,
            clients: [],
            filteredClients: [],
            searchQuery: '',
            findWindow: false,
            promo_code: '',
            verifyPromo: false,  //false - промокод не прошел проверку
            fullPromoCode: {}
        }
    },
    methods: {
        async verify_promo_code(e) {
            e.preventDefault();
            this.verifyPromo = false;

            var myHeaders = new Headers();
            myHeaders.append("Content-Type", "application/x-www-form-urlencoded");

            var urlencoded = new URLSearchParams();
            urlencoded.append("club_id", this.$props.club_id);
            urlencoded.append("promo_code", this.promo_code);

            var requestOptions = {
                method: 'POST',
                headers: myHeaders,
                body: urlencoded,
                redirect: 'follow'
            };

            var response = await fetch("api/promo/codes", requestOptions);
            if (response.ok) {
                var result = await response.json();
                this.fullPromoCode = result;
                var count = result.length;
                if (count === 0) {
                    // this.$refs.message.modal.show();
                }
                else {
                    this.verifyPromo = true;
                }
            }
            else {
                //this.$refs.message.modal.show();
            }
        },
        pay() {
            this.modal.show();
        },
        payJson() {
            this.modal.hide();
            for (var i=0; i<this.cart.length; i++){
                if (this.cart[i].allowDiscount) {
                    this.cart[i].discount = this.fullPromoCode.num == undefined ? 0 : this.fullPromoCode.num;
                    delete this.cart[i].allowDiscount;
                }
                else {
                    this.cart[i].discount = 0;
                    delete this.cart[i].allowDiscount;
                }
            }
            let json_data = JSON.stringify(this.cart);
            console.log(json_data);
            this.cart = [];
            this.fullPromoCode = {};
            this.promo_code = '';
            this.verifyPromo = false;
        },
        clearClient() {
            this.findClient = {};
            this.client = '';
        },
        searchClientToKey(e) {
            if (e.keyCode === 13) {
                this.search();
            }
        },
        searchClients() {
            this.filteredClients = this.clients.filter(client => {
                this.findWindow = this.filteredClients.length > 1;
                if (this.searchQuery.length < 2) this.findWindow = false;
                return client.Login.toLowerCase().includes(this.searchQuery.toLowerCase())
            });

        },
        handleClickOutside(event) {
            // if (!this.$refs.findBlock.contains(event.target)) {
            //     this.findWindow = false
            // }
        },
        selectFindClient(i) {
            this.searchQuery = this.filteredClients[i].login;
            setTimeout(function (){
                this.search();
            }.bind(this),200);
        },
        async getClients() {
            var myHeaders = new Headers();
            myHeaders.append("Content-Type", "application/x-www-form-urlencoded");

            var urlencoded = new URLSearchParams();
            urlencoded.append("club_id", this.$props.club_id);

            var requestOptions = {
                method: 'POST',
                headers: myHeaders,
                body: urlencoded,
                redirect: 'follow'
            };

            var response = await fetch("api/client/all", requestOptions);
            if (response.ok) {
                this.clients = await response.json();
            }
            else {
                // this.$refs.message.modal.show();
            }
        },
        async search() {
            var myHeaders = new Headers();
            myHeaders.append("Content-Type", "application/x-www-form-urlencoded");

            var urlencoded = new URLSearchParams();
            urlencoded.append("login", this.searchQuery);

            var requestOptions = {
                method: 'POST',
                headers: myHeaders,
                body: urlencoded,
                redirect: 'follow'
            };

            var response = await fetch("api/shop/find/client", requestOptions);
            this.findWindow = false;
            try {
                var result = await response.json();
                this.findClient = result;
                this.findClientStatus = true;
            }
            catch (e) {
                result = null;
            }
            if (result === null) {
                // this.$refs.sysMessage.text = 'Клиент не найден';
                // this.$refs.sysMessage.modal.show();
                this.findClient = {};
                this.findClientStatus = false;
            }
        },
        del(i) {
            this.cart.splice(i,1);
            this.calcSum();
        },
        calcSum() {
            this.rubs = 0;
            this.bonuses = 0;
            this.cart.forEach(function (item, index) {
                this.rubs += item.price * item.qty;
                this.bonuses += item.price_bonus * item.qty;
            }.bind(this))
        },
        toCart(cartQty, i, direction) {
            let warehouseQuantity = this.goods[i].Num;
            if (cartQty <= warehouseQuantity) {
                var name = this.goods[i].Product + ' ' + this.goods[i].ProductParam;
                var dublicat = false;
                this.cart.forEach(function (item, i) {
                    if (item.name === name) {
                        if (direction == 'plus')
                            this.cart[i].qty++;
                        else
                            this.cart[i].qty--;
                        dublicat = true;
                    }
                }.bind(this));
                if (!dublicat) {
                    this.cartElement.name = name;
                    this.cartElement.id = this.goods[i].storeid;
                    this.cartElement.nalog = this.goods[i].Nalog;

                    this.cartElement.allowDiscount = this.goods[i].Discount;

                    if (this.goods[i].Price != 0) {
                        this.cartElement.price = Number(this.goods[i].Price);
                        this.cartElement.price_bonus = 0;
                    }
                    else {
                        this.cartElement.price_bonus = Number(this.goods[i].PriceBonus);
                        this.cartElement.price = 0;
                    }
                    this.cartElement.qty = 1;
                    this.cart.push(this.cartElement);
                    this.cartElement = {};
                }
                this.calcSum();
            }
        },
        toCartPlus(i) {
            let cartQty = this.cart[i] == undefined ? 1 : this.cart[i].qty + 1;
            this.toCart(cartQty, i, 'plus');
        },
        toCartMinus(i) {
            let cartQty = this.cart[i] == undefined ? 1 : this.cart[i].qty - 1;
            if(cartQty > 0)
                this.toCart(cartQty, i, 'minus');
        },
        enter(e) {
            if (e.keyCode === 13) {
                this.find();
            }
        },
        async find() {
            var myHeaders = new Headers();
            myHeaders.append("Content-Type", "application/x-www-form-urlencoded");

            var urlencoded = new URLSearchParams();
            urlencoded.append("search", this.product_name);

            var requestOptions = {
                method: 'POST',
                headers: myHeaders,
                body: urlencoded,
                redirect: 'follow'
            };

            var response = await fetch("api/shop/find", requestOptions);
            if (response.ok) {
                this.goods = await response.json();
            }
            else {
                // this.$refs.message.modal.show();
            }
        },
        async getAll() {
            var myHeaders = new Headers();
            myHeaders.append("Content-Type", "application/x-www-form-urlencoded");

            var urlencoded = new URLSearchParams();
            urlencoded.append("club_id", this.$props.club_id);

            var requestOptions = {
                method: 'POST',
                headers: myHeaders,
                body: urlencoded,
                redirect: 'follow'
            };

            var response = await fetch("api/shop/get", requestOptions);
            if (response.ok) {
                var result = await response.json();
                this.goods = result[0];
                for (var i = 0; i < this.goods.length; i++) {
                    this.goods[i].Price = Number(this.goods[i].Price).toFixed(0);
                }
            }
            else {
                // this.$refs.message.modal.show();
            }
        }
    },
    mounted() {
        var payModal = document.getElementById('payModal')
        this.modal = bootstrap.Modal.getOrCreateInstance(payModal);

        this.getAll();
        this.getClients();
    }
}
</script>

<style scoped>

.pm {
    font-size: 27px;
}
.pm span {
  display: block;
  position: relative;
  top: -9px;
  left: 1px;
}
.fon {
    height: 752px;
    background: var(--light-blue-bg-color);
    border-radius: 20px;
    overflow: auto;
    width: 1491px;
    margin-left: 20px;
    margin-top: 20px;
}
.info-block{
    display: flex;
}
.button-block {
    display: flex;
    flex-direction: column;
    margin-top: 26px;
}
.user-select {
    width: 250px;
    height: 175px;

    display: flex;
    flex-direction: column;

    /*align-items: center;    */
    /*justify-content: space-between;*/
}
.bottom3 {
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 0;
    width: 250px;
    height: 212px;
    margin-top: 10px;
}
.goods-discount {
    border: 1px solid #4f0;
}
.client_login {
    display: inline-block;
    margin-top: 3px;
    margin-left: -147px;
    position: relative;
    top: 35px;
}
.modal-footer {
    justify-content: flex-start;
}
.popup {
    width: 250px;
    background: var(--light-blue-bg-color);
    position: absolute;
    top: 216px;
    left: 127px;
    border-radius: 10px;
}
ul li {
    list-style-type: none;
    cursor: pointer;
}
.bottom2 {
    overflow: auto;
    max-height: 861px;
    justify-content: space-between;
}
.save {
    background: var(--disable);
    border: none;
    margin-left: 20px;
}

.div_2 {
    height: 100px;
    display: flex;
    color: var(--standart-gray);
    flex-direction: column;
    padding-top: 10px;
}

.div_3 {
    width: 47.3%;
    height: 158px;
    display: flex;
    padding: 0 5px 0 9px;
}

.div_7 {
    width: 500px;
    height: 102px;
    display: flex;
    padding: 31px;
    padding-top: 0;
    flex-direction: column;
}

.div_4 {
    height: 35px;
    background: #172D39;
    padding: 2px 15px 5px 15px;
    margin-top: 8px;
    color: var(--standart-gray);
    border: none;
    width: 223px;
    border-radius: 8px;
}

.div_6 {
    height: 35px;
    background: #172D39;
    padding: 2px 15px 5px 15px;
    margin-top: 6px;
    color: var(--standart-gray);
    border: none;
    width: 455px;
    border-radius: 8px;
}

.check-1 + label {
    display: inline-flex;
    align-items: center;
    user-select: none;
}

.check-1 + label::after {
    content: '';
    display: inline-block;
    width: 1em;
    height: 1em;
    flex-shrink: 0;
    flex-grow: 0;
    border: 1px solid #adb5bd;
    border-radius: 0.25em;
    margin-right: 0.5em;
    background-repeat: no-repeat;
    background-position: center center;
    background-size: 50% 50%;
    left: 186px;
    position: absolute;
}

.check-1:checked + label::after {
    border-color: #0b76ef;
    background-color: #0b76ef;
    background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 8 8'%3e%3cpath fill='%23fff' d='M6.564.75l-3.59 3.612-1.538-1.55L0 4.26 2.974 7.25 8 2.193z'/%3e%3c/svg%3e");
}

.check-2 + label {
    display: inline-flex;
    align-items: center;
    user-select: none;
}

.check-2 + label::after {
    content: '';
    display: inline-block;
    width: 1em;
    height: 1em;
    flex-shrink: 0;
    flex-grow: 0;
    border: 1px solid #adb5bd;
    border-radius: 0.25em;
    margin-right: 0.5em;
    background-repeat: no-repeat;
    background-position: center center;
    background-size: 50% 50%;
    left: 416px;
    position: absolute;
}

.check-2:checked + label::after {
    border-color: #0b76ef;
    background-color: #0b76ef;
    background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 8 8'%3e%3cpath fill='%23fff' d='M6.564.75l-3.59 3.612-1.538-1.55L0 4.26 2.974 7.25 8 2.193z'/%3e%3c/svg%3e");
}

.span_1 {
    margin-top: 14px;
    display: inline-block;
    font-size: 11px;
    color: var(--standart-gray);
}

.div_data {
    display: flex;
    font-size: 13px;
    padding-left: 12px;
}

.bottom {
    height: 92px;
    display: flex;
    flex-direction: column;
    position: relative;
    top: 9px;
    left: -10px;
}

.btn {
}

.formcheck_1 {
    position: absolute;
    z-index: -1;
    opacity: 0;
}
.main2 {
    width: 500px;
    height: 200px;
}
.modal-body {
    padding: 0;
}
.content {
    background: var(--dark-green-b);
    color: var(--standart-gray);
}
.formcheck_1 + label::before {
    content: '';
    display: inline-block;
    width: 1em;
    height: 1em;
    flex-shrink: 0;
    flex-grow: 0;
    border: 1px solid #adb5bd;
    border-radius: 0.25em;
    margin-right: 0.5em;
    background-repeat: no-repeat;
    background-position: center center;
    background-size: 50% 50%;
}

.formcheck_1:checked + label::before {
    border-color: #0b76ef;
    background-color: #0b76ef;
    background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 8 8'%3e%3cpath fill='%23fff' d='M6.564.75l-3.59 3.612-1.538-1.55L0 4.26 2.974 7.25 8 2.193z'/%3e%3c/svg%3e");
}

.formcheck_1 + label::before {
    position: relative;
    left: 205px;
    top: 2px;
}

.formcheck_1 + label {
    position: relative;
    left: -23px;
    top: 5px;
}

/* стили при наведении курсора на checkbox */

.formcheck_1:not(:disabled):not(:checked) + label:hover::before {
    border-color: #b3d7ff;
}

/* стили для активного состояния чекбокса (при нажатии на него) */

.formcheck_1:not(:disabled):active + label::before {
    background-color: #b3d7ff;
    border-color: #b3d7ff;
}

/* стили для чекбокса, находящегося в фокусе */

.formcheck_1:focus + label::before {
    box-shadow: 0 0 0 0.2rem rgba(0, 123, 255, 0.25);
}

/* стили для чекбокса, находящегося в фокусе и не находящегося в состоянии checked */

.formcheck_1:focus:not(:checked) + label::before {
    border-color: #80bdff;
}

/* стили для чекбокса, находящегося в состоянии disabled */

.formcheck_1:disabled + label::before {
    background-color: #e9ecef;
}
.main {
    width: 1790px;
    height: 815px;
    display: flex;
    color: white;
    margin-left: 13px;
}

.left {
    width: 330px;
    height: 815px;
    display: flex;
    flex-direction: column;
    align-items: center;
    padding-top: 40px;
    background: var(--light-blue-bg-color);
    border-radius: 20px;
    /*justify-content: space-between;*/
    padding-bottom: 43px;
}

.right {
    width: 100%;
}

.top {
    width: 1490px;
    height: 42px;
    margin-left: 20px;
    display: flex;
    justify-content: start;
    align-items: center;
}

.input_shop {
    margin-left: 20px;
    background: var(--standart-black);
    height: 42px;
    display: flex;
    flex-direction: row;
    align-items: center;
    padding: 10px 20px;
    gap: 10px;
    width: 376px;
    box-shadow: 0px 1px 10px rgba(0, 0, 0, 0.12);
    border-radius: 10px;
    flex: none;
    order: 0;
    flex-grow: 0;
    border: none;
    color: var(--standart-color);
    padding-left: 50px;
}

.lupa {
    position: relative;
    left: 48px;
    opacity: 0.50;
    top: 0;
}

.butt {
    margin-left: 10px;
    min-width: 94px;
}

.user_name {
    width: 250px;
    height: 42px;
    background: var(--standart-black);
    border-radius: 7px;
    border: none;
    color: var(--standart-gray);
    margin-top: 20px;
}

.info {
    width: 250px;
    height: 42px;
}

.cart_element {
    width: 230px;
    min-height: 40px;
    background: var(--light-blue-bg-color);
    border-radius: 10px;
    margin-top: 19px;
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
}

.select_user {
    background: var(--standart-black);
    height: 42px;
    display: flex;
    flex-direction: row;
    align-items: center;
    padding: 10px 20px;
    gap: 10px;
    width: 250px;
    box-shadow: 0px 1px 10px rgba(0, 0, 0, 0.12);
    border-radius: 10px;
    flex: none;
    order: 0;
    flex-grow: 0;
    border: none;
    color: var(--standart-color);
    padding-left: 50px;
}

.lupa2 {
    position: relative;
    top: 31px;
    left: 10px;
}

.rub_summa {
    display: inline-block;
    position: relative;
    top: 7px;
    left: 5px;
}

.cart_name {
    width: 165px;
}

.cart_name span {
    font-size: 12px;
    line-height: 1.2;
    display: inline-block;
    margin-left: 13px;
    margin-top: 4px;
}

.cart {
    display: flex;
    width: 250px;
    height: 189px;
    flex-direction: column;
    align-items: center;
    background: var(--dark-blue-bg-color);
    border-radius: 10px;
    overflow: auto;
    margin-top: 33px;
}

::-webkit-scrollbar-track {
    -webkit-box-shadow: inset 0 0 6px rgba(0,0,0,0.3);
    border-radius: 10px;
    background-color: #142A36;
}

::-webkit-scrollbar {
    width: 12px;
    background-color: #142A36;
}

::-webkit-scrollbar-thumb {
    border-radius: 10px;
    -webkit-box-shadow: inset 0 0 6px rgba(0,0,0,.3);
    background-color: var(--dark-green);
}

.goods {
    width: 180px;
    height:270px;
    margin-right: 17px;
    margin-bottom: 17px;
    display: inline-block;
    background: var(--dark-blue-bg-color);
    padding: 10px;
    border-radius: 10px;
}

.goods_img {
    width: 160px;
    height: 170px;
    display: flex;
    justify-content: center;
    overflow: hidden;
    border-radius: 10px;
}

.goods_text {
    width: 156px;
    height: 60px;
    font-size: 10px;
    line-height: 15px;
    margin-top: 7px;
}

.goods_cart {
    border: none;
    border-radius: 22px;
    background: var(--dark-green);
    color: var(--standart-color);
    width: 22px;
    height: 22px;
}
</style>
