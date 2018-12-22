<template>
  <span><slot></slot></span>
</template>
<script>
const originalTitle = document.title;

function setup() {
  document.title = this.$el.innerText;
  if (!this.observer) {
    this.observer = new MutationObserver((/* mutations */) => {
      document.title = this.$el.innerText;
    });
    this.observer.observe(this.$el, {
      characterData: true,
      childList: true,
      subtree: true,
    });
  }
}

function teardown() {
  if (this.observer) {
    document.title = originalTitle;
    this.observer.disconnect();
    delete this.observer;
  }
}

export default {
  mounted: setup,
  activated: setup,
  deactivated: teardown,
  beforeDestroy: teardown,
};
</script>
