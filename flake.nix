{
  description = "advent-of-code shell development tooling";

  inputs.nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";
  inputs.jsonnet-tool.url = "github:jdbaldry/jsonnet-tool";

  outputs = { self, nixpkgs, flake-utils, jsonnet-tool }:
    (flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs
          {
            inherit system; overlays = [ jsonnet-tool.overlay ];
          };
      in
      {
        devShell = import ./shell.nix { inherit pkgs; };
      }));
}
