function updateLastUpdated() {
  document.querySelectorAll('.portainer-last-updated').forEach(el => {
    if (el.dataset.processed) return;

    const ts = Number(el.dataset.ts);
    if (!ts) return;

    el.textContent = new Date(ts * 1000).toLocaleString();
    el.dataset.processed = "true";
  });
}

// Run once
updateLastUpdated();

// Watch for Glance rendering / refreshes
const observer = new MutationObserver(() => {
  updateLastUpdated();
});

observer.observe(document.body, {
  childList: true,
  subtree: true
});
