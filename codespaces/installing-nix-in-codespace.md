# Installing Nixpkgs in a Codespace

## DO NOT use the installer from Determinate Systems

When using the [Determinate Systems installer](https://github.com/DeterminateSystems/nix-installer) with the linux template and without daemon, as:
```bash 
curl --proto '=https' --tlsv1.2 -sSf -L https://install.determinate.systems/nix | sh -s -- install linux --no-confirm --init none
```
This will not work properly and leads to weird problems in non nix programs like my Atuin setup. 

## Use the standard single user NixOS based installer

Using the NixOS based installer works properly:
```bash
sh <(curl -L https://nixos.org/nix/install) --no-daemon

# Configure Flakes
mkdir -p $HOME/.config/nix 2>/dev/null
echo "experimental-features = nix-command flakes" > $HOME/.config/nix/nix.conf
```

Everytime a codespace starts it is needed to fix the /tmp ACL settings to not get aborted builds as Nix expects certain chmod values:
- See https://github.com/NixOS/nix/issues/6680#issuecomment-1230902525

Fix this with:
```bash
setfacl -k /tmp
```
As this is something that can be downloaded from the cache, it is also possible to run this with Nix itself after installing:
```bash
nix shell 'nixpkgs#acl' --command bash -c 'sudo env PATH=$PATH setfacl -k /tmp'
```
