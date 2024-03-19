{
  outputs = { nixpkgs, ... }: {
    devShells.x86_64-linux.default = let
      pkgs = import nixpkgs {
        system = "x86_64-linux";
  	    config.allowUnfree = true;
      };
    in
      pkgs.mkShell {
        name = "SCTI-Web";
        packages = with pkgs; [
          go
        ];
      };
  };
}
