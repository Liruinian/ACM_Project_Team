<script>
export default {
  data() {
  	return {
      array:['a', 'b', 'c'],
  	}
	}
}
</script>

<template>
<Header>
<template v-slot:ptitle>Test</template>
</Header>
  <main>
    <TheWelcome />
  </main>
</template>

<style>

</style>
