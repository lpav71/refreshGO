<template>
  <!-- Модальное окно -->
  <div class="modal fade" id="confirmModal" tabindex="-1" aria-labelledby="confirmModalLabel" aria-hidden="true">
    <div class="modal-dialog modal-sm modal-dialog-centered">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="confirmModalLabel">{{ header }}</h5>
          <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Закрыть"></button>
        </div>
        <div class="modal-body">
          {{ text }}
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-success" data-bs-dismiss="modal" @click="yes">Да</button>
          <button type="button" class="btn btn-danger" data-bs-dismiss="modal" @click="no">Нет</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Confirm",
  data() {
    return {
      confirmModal: null,
      header: 'Подтвердите действие',
      text: 'Вы действительно хотите выполнить данную операцию?',
      result: false
    }
  },
  methods: {
    yes() {
      this.result = true;
      this.$emit('confirm', this.result);
    },
    no() {
      this.result = false;
      this.$emit('confirm', this.result);
    }
  },
  mounted() {
    var confirmModal = document.getElementById('confirmModal')
    this.confirmModal = bootstrap.Modal.getOrCreateInstance(confirmModal);
  }
}
</script>

<style scoped>

</style>