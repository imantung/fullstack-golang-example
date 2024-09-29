import Alpine from 'alpinejs';

declare global {
  interface Window {
    Alpine: typeof Alpine;
  }
}

document.addEventListener('alpine:init', () => {
  console.log("ok"); // TODO: init event
})

Alpine.start()


