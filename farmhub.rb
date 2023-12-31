# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Farmhub < Formula
  desc ""
  homepage "https://farmhub.ag"
  version "1.2.0"

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/thefarmhub/farmhub-cli/releases/download/v1.2.0/farmhub_Darwin_x86_64.tar.gz"
      sha256 "b3ed54238e450e05a4ffc3ce5493d5628becb489be29f8435186d0e935e5d39c"

      def install
        bin.install "farmhub"
      end
    end
    if Hardware::CPU.arm?
      url "https://github.com/thefarmhub/farmhub-cli/releases/download/v1.2.0/farmhub_Darwin_arm64.tar.gz"
      sha256 "f4b9db54f859293e200870298c6d3875881ca786130fbfe4a974360b697b1dda"

      def install
        bin.install "farmhub"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/thefarmhub/farmhub-cli/releases/download/v1.2.0/farmhub_Linux_arm64.tar.gz"
      sha256 "5481ffc14ab19047c1351b9204ce3d0795224c7034e2a7c5509087c70ca6ba9f"

      def install
        bin.install "farmhub"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/thefarmhub/farmhub-cli/releases/download/v1.2.0/farmhub_Linux_x86_64.tar.gz"
      sha256 "43c09dd02d030a23f62519b8e9a08ff5bb99c033436895c42adb30fc1d4d69d1"

      def install
        bin.install "farmhub"
      end
    end
  end

  test do
    system "#{bin}/farmhub version"
  end
end
