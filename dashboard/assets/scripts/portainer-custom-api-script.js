document.addEventListener("DOMContentLoaded", () => {
  // Make "Last Updated" for Portainer parse the Unix Timestamp to an actual readable date
  document.querySelectorAll('.portainer-last-updated').forEach(el => {
    const ts = Number(el.dataset.ts);
    if (!ts) return;

    el.textContent = new Date(ts * 1000).toLocaleString();
  });
});