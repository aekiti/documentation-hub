<template>
  <div class="account">
    <no-ssr>
      <AppIdenticon
        :address="value"
        class="account-identicon"
      />
    </no-ssr>
    <AppDefinition
      :title="title"
    >
      <div class="account-content">
        <FormatAddress
          :value="value"
          :icon="icon"
          :length="length"
        />
      </div>
    </AppDefinition>
  </div>
</template>
<script>
import AppDefinition from '../appDefinition'
import FormatAddress from '../formatAddress'
import AppIdenticon from '../appIdenticon'

export default {
  name: 'Account',
  components: { AppDefinition, FormatAddress, AppIdenticon },
  props: {
    value: {
      type: String,
      required: true
    },
    title: {
      type: String,
      required: true
    },
    icon: {
      type: Boolean,
      default: false
    },
    length: {
      type: String,
      default: 'responsive'
    }
  }
}
</script>
<style scoped lang="scss">
    .account {
      display: flex;
      justify-content: center;
      align-items: center;
      padding-left: .6rem;
      @media (max-width:450px) {
        flex-direction: column;
        justify-content: start;
        align-items: start;
        padding-left: 0;
      }
      &-identicon {
        @media (max-width:450px) {
          margin-left: .6rem;
          margin-top: .3rem;
        }
      }
    }

</style>
