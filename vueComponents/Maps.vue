<template>
  <div class="home">
    <div class="header_bottom">
      <div class="header_bottom_left"><i class="fas fa-info-circle"></i><span>Сведения о системе</span>
        <button class="btn btn-success" type="button">Произвести сканирование</button>
      </div>
    </div>
    <div class="main_table" style="margin-top: 20px;">
      <div class="t1">
        <table>
          <thead>
          <tr>
            <th>Номер ПК</th>
            <th>Статус ПК</th>
            <th>MAC Устройства</th>
            <th>CPU/Температура</th>
            <th>RAM</th>
            <th>GPU/Температура</th>
            <th>Материнскаяя плата</th>
            <th style="border-right-style: none;">Требуется внимание</th>
          </tr>
          </thead>
          <tbody></tbody>
        </table>
      </div>
      <div class="t2">
        <div class="scroll_table">
          <table>
            <tbody>
            <tr v-for="c in comps">
              <td>{{ c.IDComp }}</td>
              <td :style="{ color: c.status_color }">{{ c.status_ping }}</td>
              <td>{{ c.MAC }}</td>
              <td>{{ c.CPU }}</td>
              <td>{{ c.RAM }}</td>
              <td>{{ c.GPU }}</td>
              <td>{{ c.MB }}</td>
              <td style="border-right-style: none;color: green">Нет</td>
            </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  props: ['club_id'],
  data() {
    return {
      comps: []
    }
  },
  methods: {
    async getComputers() {
      const myHeaders = new Headers();
      myHeaders.append("Content-Type", "application/x-www-form-urlencoded");

      var urlencoded = new URLSearchParams();
      urlencoded.append("club_id", this.$props.club_id);

      const requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: urlencoded,
        redirect: 'follow'
      };

      const response = await fetch("http://localhost:8185/api/computer/all", requestOptions);
        if(response.ok) {
          this.comps = await response.json()
          this.comps.forEach(function (item, index) {
            if (item.ping == 1) {
              this.comps[index].status_ping = 'Доступен';
              this.comps[index].status_color = 'green';
            } else {
              this.comps[index].status_ping = 'Не доступен';
              this.comps[index].status_color = '#cc5a5a';
            }
          }.bind(this))
        }
    }
  },
  mounted() {
    this.getComputers();
  }
}
</script>

<style scoped>
.scroll_table {
  overflow-x: auto;
  height: 702px;
  background: var(--light-blue-bg-color);
}

.home {
  width: 1793px;
  height: 815px;
  color: var(--standart-gray);
  margin-left: 11px;
}

.home .header_bottom .header_bottom_left {
  display: flex;
  align-items: center;
  gap: 10px;
}

.home .header_bottom {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.home .header_bottom {
  display: flex;
  align-items: center;
  gap: 10px;
}

.main_table {
  height: 755px;
}

.home table {
  width: 100%;
  color: var(--standart-gray);
  padding: 5px;
  font-size: 11px;
  border-radius: 10px;
}

.home table th {
  width: 12.5%;
  font-size: 16px;
  color: white;
  font-weight: initial;
}

.home td {
  width: 12.5%;
  font-size: 18px;
  color: white;
  font-weight: initial;
}

.home thead {
  background: var(--light-blue-bg-color);
  border-radius: 10px;
}

.t2 {
  margin-top: 13px;
  border-radius: 15px;
  overflow: hidden;
}

.t1 th {
  padding: 10px;
  border-right-width: 1px;
  border-right-style: solid;
  text-align: center;
  border-color: #4d4d4d;
}

.t2 td {
  padding: 10px;
  border-right-width: 1px;
  border-right-style: solid;
  text-align: center;
  border-color: #4d4d4d;
}

.t1 {
  border-radius: 15px;
  overflow: hidden;
}
</style>
