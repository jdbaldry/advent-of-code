{ pkgs ? import <nixpkgs> }:

with pkgs;
mkShell {
  buildInputs = [ gnumake go_1_17 gopls go-jsonnet mitscheme rlwrap ];
  shellHook = ''
    # ...
  '';
}
