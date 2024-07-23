export const languages = {
  "plain": () => Promise.resolve({ languageSupport: () => [] }),
  "css": () => import("@codemirror/lang-css").then(module => ({ languageSupport: module.css })),
  "cpp": () => import("@codemirror/lang-cpp").then(module => ({ languageSupport: module.cpp })),
  "go": () => import("@codemirror/lang-go").then(module => ({ languageSupport: module.go })),
  "html": () => import("@codemirror/lang-html").then(module => ({ languageSupport: module.html })),
  "java": () => import("@codemirror/lang-java").then(module => ({ languageSupport: module.java })),
  "javascript": () => import("@codemirror/lang-javascript").then(module => ({ languageSupport: module.javascript })),
  "json": () => import("@codemirror/lang-json").then(module => ({ languageSupport: module.json })),
  "python": () => import("@codemirror/lang-python").then(module => ({ languageSupport: module.python })),
  "rust": () => import("@codemirror/lang-rust").then(module => ({ languageSupport: module.rust })),
  "markdown": () => import("@codemirror/lang-markdown").then(module => ({ languageSupport: module.markdown })),
  "php": () => import("@codemirror/lang-php").then(module => ({ languageSupport: module.php })),
  "xml": () => import("@codemirror/lang-xml").then(module => ({ languageSupport: module.xml })),
  "yaml": () => import("@codemirror/lang-yaml").then(module => ({ languageSupport: module.yaml })),
};
