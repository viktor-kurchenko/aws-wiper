package internal

import (
	"fmt"
	"log/slog"
	"strings"
	"time"
)

type regionCluster struct {
	Region  string
	Cluster string
}

func Wipe(clusters []string, tries int, pause time.Duration) error {
	rClusters, err := parseClusters(clusters)
	if err != nil {
		return err
	}
	for i := range rClusters {
		slog.Info("deleting:", rClusters[i].Region, rClusters[i].Cluster)
	}
	//todo: implement cleanup ...
	return nil
}

func parseClusters(clusters []string) ([]regionCluster, error) {
	rClusters := make([]regionCluster, 0, len(clusters))
	for _, c := range clusters {
		parts := strings.Split(c, "=")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid region=cluster pair: %s", c)
		}
		rClusters = append(rClusters, regionCluster{
			Region:  parts[0],
			Cluster: parts[1],
		})
	}
	return rClusters, nil
}
