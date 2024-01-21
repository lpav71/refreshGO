<template>
    <meta name="csrf-token" content="{{ csrf_token() }}">
    <div class="account">
        <form action="finduser" method="get">
            <div class="left-user-account"><i class="fa fa-search fs-5 lupa"></i>
                <input type="text" name="searchUser" @input="searchClients" v-model="searchQuery" class="user-search" placeholder="Поиск пользователя">
            <div class="popup" v-if="findWindow" ref="findBlock">
                <ul>
                    <li v-for="(f, i) in filteredClients" @click="selectFindClient(i)">{{ f.login }}</li>
                </ul>
            </div>
        </div>
        <button ref="send" type="submit" style="display: none"></button>
        </form>
        <div class="right-user-account">
            <div class="user-search-button" @click="listMessageToggle"><i class="fa fa-comment fs-3 icon-user-plus"></i></div>
            <div class="user-search-button" style="padding-left: 10px;"><i class="fa fa-bell fs-3 icon-user-plus"></i>
                <div class="notification"><span class="count-notif" style="color: var(--standart-color);font-size: 9px;">10</span></div>
            </div>
            <div class="user-select-button">
                <div class="avatar"><img style="width: 42px;height: 42px;" :src="icon" alt=""></div>
                <select v-model="user_data" class="user-change" @change="promo">
                    <option value="name" selected="selected">{{ name }}</option>
                    <option value="promo">Промокоды</option>
                    <option value="dashboard">Статистика</option>
                    <option value="task">История задач</option>
                    <option value="infouser">Информация о пользователе</option>
                    <option value="tariffs">Тарифы и предложения</option>
                    <option value="pay_terminal">Настройки ККМ</option>
                    <option value="position">Сотрудники и роли</option>
                    <option value="licenses">Лицензии</option>
                    <option value="design">Дизайн</option>
                    <option value="users-group">Группы пользователей</option>
                    <option value="history-deposits">История пополнений</option>
                    <option value="history-sales">История продаж</option>
                    <option value="tariff_statistic">Статистика тарифов</option>
                    <option value="sound">Настройка звука</option>
                    <option value="receipt">Поступления</option>
                    <option value="version">Версии ПО</option>
                    <option value="customer-purchases">Покупки клиентов</option>
                    <option value="return-info">Информация по списаниям</option>
                    <option value="logout">Выход</option>
                </select>
            </div>
        </div>
    </div>

<!--    <message ref="message"></message>-->

</template>

<script>
export default {
    props: ['name', 'club_id', 'user_id', 'icon'],
    data() {
        return {
            modal: null,
            login: null,
            pass: null,
            lastName: null,  //Фамилия
            firstName: null, //Имя
            patronymic: null, //Отчество
            bday: null,
            phone: null,
            address: null,
            email: null,
            vkId: null,
            user_data: 'name',
            clients: [],
            filteredClients: [],
            searchQuery: '',
            findWindow: false,
            permissions: {},
            listMessageStatus: false,
            listMessageOpacity: 0,
            listMessageScale: 0.5,
            listMessageText: [],
            logoutForm: null
        }
    },
    methods: {
        async promo(e) {
            if (this.user_data == 'logout') {
                e.preventDefault();
                var token = document.querySelector('meta[name="csrf-token"]').getAttribute('content');
                var myHeaders = new Headers();
                myHeaders.append("Content-Type", "application/x-www-form-urlencoded");
                myHeaders.append("X-CSRF-TOKEN", token);

                var urlencoded = new URLSearchParams();

                var requestOptions = {
                    method: 'POST',
                    headers: myHeaders,
                    body: urlencoded,
                    redirect: 'follow'
                };

                var response = await fetch("logout", requestOptions);
                if (response.ok){
                    window.location.href = 'login';
                }
            }
            if (this.user_data == 'promo') {
                window.open("promo", '_self');
            }
            if (this.user_data == 'dashboard') {
                window.open("dashboard", '_self');
            }
            if (this.user_data == 'task') {
                window.open("task", '_self');
            }
            if (this.user_data == 'infouser') {
                window.open("userinfo", '_self');
            }
            if (this.user_data == 'tariffs') {
                window.open("tariffs", '_self');
            }
            if (this.user_data == 'pay_terminal') {
                window.open("pay_terminal", '_self');
            }
            if (this.user_data == 'position') {
                window.open("position", '_self');
            }
            if (this.user_data == 'licenses') {
                window.open("licenses", '_self');
            }
            if (this.user_data == 'design') {
                window.open("design", '_self');
            }
            if (this.user_data == 'users-group') {
                window.open("users-group", '_self');
            }
            if (this.user_data == 'history-deposits') {
                window.open("history-deposits", '_self');
            }
            if (this.user_data == 'history-sales') {
                window.open("history-sales", '_self');
            }
            if (this.user_data == 'tariff_statistic') {
                window.open("tariff_statistic", '_self');
            }
            if (this.user_data == 'sound') {
                window.open("sound", '_self');
            }
            if (this.user_data == 'receipt') {
                window.open("receipt", '_self');
            }
            if (this.user_data == 'version') {
                window.open("version", '_self');
            }
            if (this.user_data == 'customer-purchases') {
                window.open("customer-purchases", '_self');
            }
            if (this.user_data == 'return-info') {
                window.open("return-info", '_self');
            }
        },
        listMessageToggle() {
            this.listMessageStatus = !this.listMessageStatus;
            if (this.listMessageStatus){
                this.listMessageOpacity = 1;
                this.listMessageScale = 1;
                this.getMessage();
            }
            else {
                this.listMessageOpacity = 0;
                this.listMessageScale = 0.5;
            }
        },
        listMessageHide() {
            this.listMessageStatus = false;
            this.listMessageOpacity = 0;
            this.listMessageScale = 0.5;
        },
        selectFindClient(i) {
            this.searchQuery = this.filteredClients[i].login;
            setTimeout(function (){
                this.$refs.send.click();
            }.bind(this),200);
        },
        async getMessage() {
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

            var response = await fetch("api/message", requestOptions);
            if (response.ok) {
                this.listMessageText = await response.json();
            }
            else {
                //this.$refs.message.modal.show();
            }
        },
        async getPermissions() {
            var myHeaders = new Headers();
            myHeaders.append("Content-Type", "application/x-www-form-urlencoded");

            var urlencoded = new URLSearchParams();
            urlencoded.append("club_id", this.$props.club_id);
            urlencoded.append("user_id", this.$props.user_id);

            var requestOptions = {
                method: 'POST',
                headers: myHeaders,
                body: urlencoded,
                redirect: 'follow'
            };

            var response = await fetch("api/user/getpermissions", requestOptions);
            if (response.ok) {
                const permissions = await response.json();
                this.permissions = permissions[0];
            }
            else {
                //this.$refs.message.modal.show();
            }
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
                //this.$refs.message.modal.show();
            }
        },
        searchClients() {
            this.filteredClients = this.clients.filter(client => {
                this.findWindow = this.filteredClients.length > 1;
                if (this.searchQuery.length < 2) this.findWindow = false;
                return client.login.toLowerCase().includes(this.searchQuery.toLowerCase())
            });

        },
    },
    mounted() {
        //this.getClients();
        this.getPermissions();
        document.addEventListener('keydown', function(event) {
            if (event.key === 'Escape') {
                this.listMessageHide();
            }
        }.bind(this));
    },
}
</script>

<style scoped>

.user-search {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 10px;
    width: 376px;
    height: 54px;
    background: linear-gradient(180deg, #284F62 0%, #183F52 100%);
    box-shadow: 0 1px 10px rgba(0, 0, 0, 0.12);
    border-radius: 10px;
    flex: none;
    order: 0;
    flex-grow: 0;
    border: none;
    color: var(--standart-color);
    padding: 10px 20px 10px 50px;
    margin-left: -6px;
}
.account {
    position: relative;
    display: flex;
    flex-direction: row;
    align-items: flex-start;
    width: 1820px;
    height: 94px;
    top: 1px;
    justify-content: space-between;
    padding: 20px 15px 20px 0;
}

ul li {
    list-style-type: none;
    cursor: pointer;
}
.popup {
    width: 376px;
    background: var(--light-blue-bg-color);
    position: absolute;
    top: 80px;
    left: 12px;
    border-radius: 10px;
    z-index: 5;
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

.account {
    position: relative;
    display: flex;
    flex-direction: row;
    align-items: flex-start;
    padding: 20px 0px;
    width: 1820px;
    height: 94px;
    top: 1px;
    justify-content: space-between;
    padding-right: 15px;
}
.user-search {
    display: flex;
    flex-direction: row;
    align-items: center;
    padding: 10px 20px;
    gap: 10px;
    width: 376px;
    height: 54px;
    background: linear-gradient(180deg, #284F62 0%, #183F52 100%);
    box-shadow: 0px 1px 10px rgba(0, 0, 0, 0.12);
    border-radius: 10px;
    flex: none;
    order: 0;
    flex-grow: 0;
    border: none;
    color: var(--standart-color);
    padding-left: 50px;
    margin-left: -6px;
}
.left-user-account {
    display: flex;
    color: var(--standart-gray);
}
.lupa {
    position: relative;
    left: 30px;
    top: 19px;
    opacity: 0.5;
}
.user-search-button {
    height: 54px;
    width: 52px;
    background: var(--light-blue-bg-color);
    box-shadow: 0px 1px 10px rgba(0, 0, 0, 0.12);
    border-radius: 10px;
    margin-left: 10px;
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
}
.right-user-account {
    display: flex;
}
.icon-user-plus {
    opacity: 0.5;
    color: var(--standart-gray);
}
.user-select-button {
    width: 263px;
    height: 54px;
    background: var(--light-blue-bg-color);
    border-radius: 10px;
    margin-left: 10px;
    display: flex;
    justify-content: space-between;
    align-items: center;
}
.avatar {
    border-radius: 10px;
    overflow: hidden;
    margin-left: 10px;
}
.user-change {
    background: var(--light-blue-bg-color);
    width: 190px;
    color: var(--standart-gray);
    margin-right: 10px;
    border: none;
    height: 54px;
}
.user-change option {
    color: var(--standart-gray);
    background-color: #172d39;
    font-size: 15px;
}
.notification {
    width: 15px;
    height: 15px;
    background: var(--light-green);
    border-radius: 11px;
    position: relative;
    left: -9px;
    top: -11px;
    display: flex;
    justify-content: center;
    align-items: center;
}
</style>
