{
  description = "Base16 theme switcher with hot-reload";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      systems = [ "x86_64-linux" "aarch64-linux" ];
      forAllSystems = nixpkgs.lib.genAttrs systems;
    in
    {
      packages = forAllSystems (system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        {
          default = pkgs.buildGoModule {
            pname = "base16changer";
            version = "0.1.0";
            src = ./.;
            vendorHash = "sha256-rvJRU+asDpk/qxCbGGps2JZtAptlMAymphSu9SEzFmI=";

            meta = {
              description = "Base16 theme switcher with hot-reload for labwc, kitty, fuzzel, GTK";
              mainProgram = "base16changer";
            };
          };
        });

      devShells = forAllSystems (system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [
              go
              gopls
            ];
          };
        });
    };
}
