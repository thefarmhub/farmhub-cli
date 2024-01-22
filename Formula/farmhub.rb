# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Farmhub < Formula
  desc ""
  homepage "https://farmhub.ag"
  version "1.5.11"

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/thefarmhub/farmhub-cli/releases/download/v1.5.11/farmhub_Darwin_x86_64.tar.gz"
      sha256 "dccd8c4504067defe62a3318fa10b62da3d217d4cc4ceb6c977c294110086aad"

      def install
        bin.install "farmhub"
      end
    end
    if Hardware::CPU.arm?
      url "https://github.com/thefarmhub/farmhub-cli/releases/download/v1.5.11/farmhub_Darwin_arm64.tar.gz"
      sha256 "4280e5d159eaa2f8fde1bb16cdfad380d854ef70e6d9d0f8cb0eb960bfc5e7fa"

      def install
        bin.install "farmhub"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/thefarmhub/farmhub-cli/releases/download/v1.5.11/farmhub_Linux_arm64.tar.gz"
      sha256 "09cb09d4ae696154d298b4bf0bc0f1c552dc99b4a73088e47b96bccd9f7328dd"

      def install
        bin.install "farmhub"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/thefarmhub/farmhub-cli/releases/download/v1.5.11/farmhub_Linux_x86_64.tar.gz"
      sha256 "ebd44b82abc1c9eb91aff36dd6afbbd06aa4fd8110865d5a6e48d5b3e37f8cb0"

      def install
        bin.install "farmhub"
      end
    end
  end

  test do
    system "#{bin}/farmhub version"
  end
end
