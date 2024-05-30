{ lib, buildGoModule, fetchFromGitHub }:


buildGoModule rec {
  pname = "scratch";
  version = "1.0.2";

  src = fetchFromGitHub {
    owner = "firesquid6";
    repo = "scratch";
    rev = "v${version}";
    hash = "sha256-FNTC82AaDhH/17zPYDmBeO/jIRT1R7bOLgemxwH1Na8";
  };

  vendorSha256 = null;

  meta = with lib; {
    description = "A simple project manager";
    homepage = "https://github.com/firesquid6/scratch";
    license = licenses.gpl3;
    longDescription = ''
      Scratch is a simple project manager taht allows you to easily create and manage projects
      and scratchpads through the command line.
    '';
    maintainers = with maintainers; [ firesquid6 ];
  };
}
