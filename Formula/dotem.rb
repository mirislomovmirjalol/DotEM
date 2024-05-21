class Dotem < Formula
  desc "Manage your environment variables easily"
  homepage "https://github.com/mirislomovmirjalol/DotEM"
  url "https://github.com/mirislomovmirjalol/DotEM/archive/v1.0.0.tar.gz"
  sha256 "00cb5d7789b897c44fa83009673e4e25009afcca5bc1895b2ea3f0a1a4e10724"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", bin/".em", "."
  end

  test do
    system "#{bin}/.em", "--version"
  end
end

