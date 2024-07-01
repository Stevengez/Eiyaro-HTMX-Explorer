const t = (o) => {
  const n = window.location.href.split("/");
  n[3] = o.toLowerCase(), window.location = n.join("/");
}, e = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
  __proto__: null,
  onLangChange: t
}, Symbol.toStringTag, { value: "Module" }));
Object.keys(e).forEach((o) => {
  console.log("Key", o), window[o] = e[o];
});
