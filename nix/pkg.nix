{ lib, buildGoModule, fetchFromGitHub }:


buildGoModule rec {
  pname = "scratch";
  version = "1.0.1";

  src = fetchFromGittHub {
    owner = "firesquid6";
    repo = "scratch";
    rev = "v${version}";
    hash = ""
  };

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
