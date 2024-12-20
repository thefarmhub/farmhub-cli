# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Farmhub < Formula
  desc ""
  homepage "https://farmhub.ag"
  version "1.5.22"

  on_macos do
    on_intel do
      url "https://github.com/thefarmhub/farmhub-cli/releases/download/v1.5.22/farmhub_Darwin_x86_64.tar.gz"
      sha256 "43fa5a8aa80021c6e873fa772633bc5fbf22bd8c0cea96d9b6b80768da5efc8d"

      def install
        bin.install "farmhub"
      end
    end
    on_arm do
      url "https://github.com/thefarmhub/farmhub-cli/releases/download/v1.5.22/farmhub_Darwin_arm64.tar.gz"
      sha256 "d2cec8b6b8886bfbc66d8ae9ba3c502e179859ab1e6ab73ffd5abe6fe080e1a1"

      def install
        bin.install "farmhub"
      end
    end
  end

  on_linux do
    on_intel do
      if Hardware::CPU.is_64_bit?
        url "https://github.com/thefarmhub/farmhub-cli/releases/download/v1.5.22/farmhub_Linux_x86_64.tar.gz"
        sha256 "710b916fea92b22994cf05e0f70ebe7d186a1378f0359f4cece673c5bb0499d8"

        def install
          bin.install "farmhub"
        end
      end
    end
    on_arm do
      if Hardware::CPU.is_64_bit?
        url "https://github.com/thefarmhub/farmhub-cli/releases/download/v1.5.22/farmhub_Linux_arm64.tar.gz"
        sha256 "75b54b53768d63f535f18c4da4159cd63d367a7b427df83c7c174de6bfbde7ad"

        def install
          bin.install "farmhub"
        end
      end
    end
  end

  test do
    system "#{bin}/farmhub version"
  end
end
