{ pkgs ? import <nixpkgs> }:

with pkgs;
mkShell {
  buildInputs = [ gnumake go-jsonnet mitscheme rlwrap ];
  shellHook = ''
    # ...
  '';
}
