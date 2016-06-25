package main

type Weights []*Weight

func (ws Weights) Len() int {
	return len(ws)
}

func (ws Weights) Less(i, j int) bool {
	return ws[i].Value > ws[j].Value
}

func (ws Weights) Swap(i, j int) {
	ws[i], ws[j] = ws[j], ws[i]
}

func weights(s []string) (Weights, error) {
	ws := make(Weights, 0, len(s))
	for _, s := range s {
		w, err := weight(s)
		if err != nil {
			return nil, err
		}
		ws = append(ws, w)
	}
	return ws, nil
}
