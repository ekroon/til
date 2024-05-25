# Hide app in application switcher

To hide an app in the MacOS application switcher (and Dock), add the following to `Info.plist`[^1]:

```xml
<key>LSUIElement</key>
<string>1</string>
```

[^1]: https://developer.apple.com/documentation/bundleresources/information_property_list/lsuielement
