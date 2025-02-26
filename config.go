package goffmpeg

import (
	"context"
	"errors"
	"runtime"
	"strings"

	"github.com/VictoriqueMoe/goffmpeg/pkg/cmd"
)

const (
	ffmpegCommand  = "ffmpeg"
	ffprobeCommand = "ffprobe"
)

type Configuration struct {
	FfprobeBinPath string
	FfmpegBinPath  string
}

func (cfg Configuration) FFmpegBinPath() string {
	return cfg.FfmpegBinPath
}

func (cfg Configuration) FFprobeBinPath() string {
	return cfg.FfprobeBinPath
}

func NewConfiguration(ffmpegBinPath, ffprobeBinPath string) Configuration {
	return Configuration{
		FfmpegBinPath:  normalizeBinPath(ffmpegBinPath),
		FfprobeBinPath: normalizeBinPath(ffprobeBinPath),
	}
}

func Configure(ctx context.Context) (Configuration, error) {
	ffmpegBin, err := cmd.FindBinPath(ctx, ffmpegCommand)
	if err != nil {
		return Configuration{}, err
	}

	if ffmpegBin == "" {
		return Configuration{}, errors.New("ffmpeg not found, please install it before using goffmpeg")
	}

	ffprobeBin, err := cmd.FindBinPath(ctx, ffprobeCommand)
	if err != nil {
		return Configuration{}, err
	}

	if ffprobeBin == "" {
		return Configuration{}, errors.New("ffprobe not found, please install it before using goffmpeg")
	}

	return Configuration{
		FfmpegBinPath:  normalizeBinPath(ffmpegBin),
		FfprobeBinPath: normalizeBinPath(ffprobeBin),
	}, nil
}

func normalizeBinPath(binPath string) string {
	binPath = strings.ReplaceAll(binPath, lineSeparator(), " ")
	return strings.TrimSpace(binPath)
}

func lineSeparator() string {
	switch runtime.GOOS {
	case "windows":
		return "\r\n"
	default:
		return "\n"
	}
}
