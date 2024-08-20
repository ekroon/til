# LlamaLife minimize width with tamper monkey (275px)

To fit Llama (with Rectangle Pro's Pin mode) in a 275px width, use this Tamper Monkey script:

```javascript
// ==UserScript==
// @name         LlamaLife
// @namespace    http://tampermonkey.net/
// @version      2024-05-23
// @description  try to take over the world!
// @author       You
// @match        https://llamalife.co/tasks
// @icon         https://www.google.com/s2/favicons?sz=64&domain=llamalife.co
// @grant        GM_addStyle
// ==/UserScript==

(function() {
    'use strict';
    GM_addStyle('.navbar > div > div:nth-child(3) { display: none }');
    // GM_addStyle('::-webkit-scrollbar{width:0px;height:3px}::-webkit-scrollbar-thumb{background-color:#c1c1c1;}::-webkit-scrollbar-thumb:hover{background-color:#666}');
    GM_addStyle('::-webkit-scrollbar{width:0px}');
    GM_addStyle('#root > div > div > div:nth-child(1) { padding-left: 0; padding-right: 0 }');
    GM_addStyle('#root > div > div > div > div:nth-child(2) { margin-left: 0 !important; }');
})();
```
