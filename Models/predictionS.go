package models

import (
	"time"

	"gorm.io/gorm"
)

type PredictionAPI struct {
	Get        string `json:"get"`
	Parameters struct {
		Fixture string `json:"fixture"`
	} `json:"parameters"`
	Errors  []interface{} `json:"errors"`
	Results int           `json:"results"`
	Paging  struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []struct {
		Predictions struct {
			Winner struct {
				ID      int    `json:"id"`
				Name    string `json:"name"`
				Comment string `json:"comment"`
			} `json:"winner"`
			WinOrDraw bool   `json:"win_or_draw"`
			UnderOver string `json:"under_over"`
			Goals     struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"goals"`
			Advice  string `json:"advice"`
			Percent struct {
				Home string `json:"home"`
				Draw string `json:"draw"`
				Away string `json:"away"`
			} `json:"percent"`
		} `json:"predictions"`
		League struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Country string `json:"country"`
			Logo    string `json:"logo"`
			Flag    string `json:"flag"`
			Season  int    `json:"season"`
		} `json:"league"`
		Teams struct {
			Home struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Logo  string `json:"logo"`
				Last5 struct {
					Form  string `json:"form"`
					Att   string `json:"att"`
					Def   string `json:"def"`
					Goals struct {
						For struct {
							Total   int    `json:"total"`
							Average string `json:"average"`
						} `json:"for"`
						Against struct {
							Total   int    `json:"total"`
							Average string `json:"average"`
						} `json:"against"`
					} `json:"goals"`
				} `json:"last_5"`
				League struct {
					Form     string `json:"form"`
					Fixtures struct {
						Played struct {
							Home  int `json:"home"`
							Away  int `json:"away"`
							Total int `json:"total"`
						} `json:"played"`
						Wins struct {
							Home  int `json:"home"`
							Away  int `json:"away"`
							Total int `json:"total"`
						} `json:"wins"`
						Draws struct {
							Home  int `json:"home"`
							Away  int `json:"away"`
							Total int `json:"total"`
						} `json:"draws"`
						Loses struct {
							Home  int `json:"home"`
							Away  int `json:"away"`
							Total int `json:"total"`
						} `json:"loses"`
					} `json:"fixtures"`
					Goals struct {
						For struct {
							Total struct {
								Home  int `json:"home"`
								Away  int `json:"away"`
								Total int `json:"total"`
							} `json:"total"`
							Average struct {
								Home  string `json:"home"`
								Away  string `json:"away"`
								Total string `json:"total"`
							} `json:"average"`
							Minute struct {
								Zero15 struct {
									Total      interface{} `json:"total"`
									Percentage interface{} `json:"percentage"`
								} `json:"0-15"`
								One630 struct {
									Total      int    `json:"total"`
									Percentage string `json:"percentage"`
								} `json:"16-30"`
								Three145 struct {
									Total      int    `json:"total"`
									Percentage string `json:"percentage"`
								} `json:"31-45"`
								Four660 struct {
									Total      interface{} `json:"total"`
									Percentage interface{} `json:"percentage"`
								} `json:"46-60"`
								Six175 struct {
									Total      int    `json:"total"`
									Percentage string `json:"percentage"`
								} `json:"61-75"`
								Seven690 struct {
									Total      int    `json:"total"`
									Percentage string `json:"percentage"`
								} `json:"76-90"`
								Nine1105 struct {
									Total      interface{} `json:"total"`
									Percentage interface{} `json:"percentage"`
								} `json:"91-105"`
								One06120 struct {
									Total      interface{} `json:"total"`
									Percentage interface{} `json:"percentage"`
								} `json:"106-120"`
							} `json:"minute"`
						} `json:"for"`
						Against struct {
							Total struct {
								Home  int `json:"home"`
								Away  int `json:"away"`
								Total int `json:"total"`
							} `json:"total"`
							Average struct {
								Home  string `json:"home"`
								Away  string `json:"away"`
								Total string `json:"total"`
							} `json:"average"`
							Minute struct {
								Zero15 struct {
									Total      int    `json:"total"`
									Percentage string `json:"percentage"`
								} `json:"0-15"`
								One630 struct {
									Total      int    `json:"total"`
									Percentage string `json:"percentage"`
								} `json:"16-30"`
								Three145 struct {
									Total      int    `json:"total"`
									Percentage string `json:"percentage"`
								} `json:"31-45"`
								Four660 struct {
									Total      interface{} `json:"total"`
									Percentage interface{} `json:"percentage"`
								} `json:"46-60"`
								Six175 struct {
									Total      int    `json:"total"`
									Percentage string `json:"percentage"`
								} `json:"61-75"`
								Seven690 struct {
									Total      int    `json:"total"`
									Percentage string `json:"percentage"`
								} `json:"76-90"`
								Nine1105 struct {
									Total      int    `json:"total"`
									Percentage string `json:"percentage"`
								} `json:"91-105"`
								One06120 struct {
									Total      interface{} `json:"total"`
									Percentage interface{} `json:"percentage"`
								} `json:"106-120"`
							} `json:"minute"`
						} `json:"against"`
					} `json:"goals"`
					Biggest struct {
						Streak struct {
							Wins  int `json:"wins"`
							Draws int `json:"draws"`
							Loses int `json:"loses"`
						} `json:"streak"`
						Wins struct {
							Home string `json:"home"`
							Away string `json:"away"`
						} `json:"wins"`
						Loses struct {
							Home string `json:"home"`
							Away string `json:"away"`
						} `json:"loses"`
						Goals struct {
							For struct {
								Home int `json:"home"`
								Away int `json:"away"`
							} `json:"for"`
							Against struct {
								Home int `json:"home"`
								Away int `json:"away"`
							} `json:"against"`
						} `json:"goals"`
					} `json:"biggest"`
					CleanSheet struct {
						Home  int `json:"home"`
						Away  int `json:"away"`
						Total int `json:"total"`
					} `json:"clean_sheet"`
					FailedToScore struct {
						Home  int `json:"home"`
						Away  int `json:"away"`
						Total int `json:"total"`
					} `json:"failed_to_score"`
					Penalty struct {
						Scored struct {
							Total      int    `json:"total"`
							Percentage string `json:"percentage"`
						} `json:"scored"`
						Missed struct {
							Total      int    `json:"total"`
							Percentage string `json:"percentage"`
						} `json:"missed"`
						Total int `json:"total"`
					} `json:"penalty"`
					Lineups []interface{} `json:"lineups"`
					Cards   struct {
						Yellow struct {
							Zero15 struct {
								Total      int    `json:"total"`
								Percentage string `json:"percentage"`
							} `json:"0-15"`
							One630 struct {
								Total      int    `json:"total"`
								Percentage string `json:"percentage"`
							} `json:"16-30"`
							Three145 struct {
								Total      int    `json:"total"`
								Percentage string `json:"percentage"`
							} `json:"31-45"`
							Four660 struct {
								Total      int    `json:"total"`
								Percentage string `json:"percentage"`
							} `json:"46-60"`
							Six175 struct {
								Total      int    `json:"total"`
								Percentage string `json:"percentage"`
							} `json:"61-75"`
							Seven690 struct {
								Total      int    `json:"total"`
								Percentage string `json:"percentage"`
							} `json:"76-90"`
							Nine1105 struct {
								Total      interface{} `json:"total"`
								Percentage interface{} `json:"percentage"`
							} `json:"91-105"`
							One06120 struct {
								Total      interface{} `json:"total"`
								Percentage interface{} `json:"percentage"`
							} `json:"106-120"`
						} `json:"yellow"`
						Red struct {
							Zero15 struct {
								Total      interface{} `json:"total"`
								Percentage interface{} `json:"percentage"`
							} `json:"0-15"`
							One630 struct {
								Total      interface{} `json:"total"`
								Percentage interface{} `json:"percentage"`
							} `json:"16-30"`
							Three145 struct {
								Total      interface{} `json:"total"`
								Percentage interface{} `json:"percentage"`
							} `json:"31-45"`
							Four660 struct {
								Total      interface{} `json:"total"`
								Percentage interface{} `json:"percentage"`
							} `json:"46-60"`
							Six175 struct {
								Total      interface{} `json:"total"`
								Percentage interface{} `json:"percentage"`
							} `json:"61-75"`
							Seven690 struct {
								Total      interface{} `json:"total"`
								Percentage interface{} `json:"percentage"`
							} `json:"76-90"`
							Nine1105 struct {
								Total      interface{} `json:"total"`
								Percentage interface{} `json:"percentage"`
							} `json:"91-105"`
							One06120 struct {
								Total      interface{} `json:"total"`
								Percentage interface{} `json:"percentage"`
							} `json:"106-120"`
						} `json:"red"`
					} `json:"cards"`
				} `json:"league"`
			} `json:"home"`
			Away struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Logo  string `json:"logo"`
				Last5 struct {
					Form  string `json:"form"`
					Att   string `json:"att"`
					Def   string `json:"def"`
					Goals struct {
						For struct {
							Total   int    `json:"total"`
							Average string `json:"average"`
						} `json:"for"`
						Against struct {
							Total   int    `json:"total"`
							Average string `json:"average"`
						} `json:"against"`
					} `json:"goals"`
				} `json:"last_5"`
				League struct {
					Form     string `json:"form"`
					Fixtures struct {
						Played struct {
							Home  int `json:"home"`
							Away  int `json:"away"`
							Total int `json:"total"`
						} `json:"played"`
						Wins struct {
							Home  int `json:"home"`
							Away  int `json:"away"`
							Total int `json:"total"`
						} `json:"wins"`
						Draws struct {
							Home  int `json:"home"`
							Away  int `json:"away"`
							Total int `json:"total"`
						} `json:"draws"`
						Loses struct {
							Home  int `json:"home"`
							Away  int `json:"away"`
							Total int `json:"total"`
						} `json:"loses"`
					} `json:"fixtures"`
					Goals struct {
						For struct {
							Total struct {
								Home  int `json:"home"`
								Away  int `json:"away"`
								Total int `json:"total"`
							} `json:"total"`
							Average struct {
								Home  string `json:"home"`
								Away  string `json:"away"`
								Total string `json:"total"`
							} `json:"average"`
							Minute struct {
								Zero15 struct {
									Total      int    `json:"total"`
									Percentage string `json:"percentage"`
								} `json:"0-15"`
								One630 struct {
									Total      interface{} `json:"total"`
									Percentage interface{} `json:"percentage"`
								} `json:"16-30"`
								Three145 struct {
									Total      int    `json:"total"`
									Percentage string `json:"percentage"`
								} `json:"31-45"`
								Four660 struct {
									Total      int    `json:"total"`
									Percentage string `json:"percentage"`
								} `json:"46-60"`
								Six175 struct {
									Total      int    `json:"total"`
									Percentage string `json:"percentage"`
								} `json:"61-75"`
								Seven690 struct {
									Total      int    `json:"total"`
									Percentage string `json:"percentage"`
								} `json:"76-90"`
								Nine1105 struct {
									Total      int    `json:"total"`
									Percentage string `json:"percentage"`
								} `json:"91-105"`
								One06120 struct {
									Total      interface{} `json:"total"`
									Percentage interface{} `json:"percentage"`
								} `json:"106-120"`
							} `json:"minute"`
						} `json:"for"`
						Against struct {
							Total struct {
								Home  int `json:"home"`
								Away  int `json:"away"`
								Total int `json:"total"`
							} `json:"total"`
							Average struct {
								Home  string `json:"home"`
								Away  string `json:"away"`
								Total string `json:"total"`
							} `json:"average"`
							Minute struct {
								Zero15 struct {
									Total      int    `json:"total"`
									Percentage string `json:"percentage"`
								} `json:"0-15"`
								One630 struct {
									Total      int    `json:"total"`
									Percentage string `json:"percentage"`
								} `json:"16-30"`
								Three145 struct {
									Total      int    `json:"total"`
									Percentage string `json:"percentage"`
								} `json:"31-45"`
								Four660 struct {
									Total      int    `json:"total"`
									Percentage string `json:"percentage"`
								} `json:"46-60"`
								Six175 struct {
									Total      int    `json:"total"`
									Percentage string `json:"percentage"`
								} `json:"61-75"`
								Seven690 struct {
									Total      int    `json:"total"`
									Percentage string `json:"percentage"`
								} `json:"76-90"`
								Nine1105 struct {
									Total      int    `json:"total"`
									Percentage string `json:"percentage"`
								} `json:"91-105"`
								One06120 struct {
									Total      interface{} `json:"total"`
									Percentage interface{} `json:"percentage"`
								} `json:"106-120"`
							} `json:"minute"`
						} `json:"against"`
					} `json:"goals"`
					Biggest struct {
						Streak struct {
							Wins  int `json:"wins"`
							Draws int `json:"draws"`
							Loses int `json:"loses"`
						} `json:"streak"`
						Wins struct {
							Home string `json:"home"`
							Away string `json:"away"`
						} `json:"wins"`
						Loses struct {
							Home string `json:"home"`
							Away string `json:"away"`
						} `json:"loses"`
						Goals struct {
							For struct {
								Home int `json:"home"`
								Away int `json:"away"`
							} `json:"for"`
							Against struct {
								Home int `json:"home"`
								Away int `json:"away"`
							} `json:"against"`
						} `json:"goals"`
					} `json:"biggest"`
					CleanSheet struct {
						Home  int `json:"home"`
						Away  int `json:"away"`
						Total int `json:"total"`
					} `json:"clean_sheet"`
					FailedToScore struct {
						Home  int `json:"home"`
						Away  int `json:"away"`
						Total int `json:"total"`
					} `json:"failed_to_score"`
					Penalty struct {
						Scored struct {
							Total      int    `json:"total"`
							Percentage string `json:"percentage"`
						} `json:"scored"`
						Missed struct {
							Total      int    `json:"total"`
							Percentage string `json:"percentage"`
						} `json:"missed"`
						Total int `json:"total"`
					} `json:"penalty"`
					Lineups []interface{} `json:"lineups"`
					Cards   struct {
						Yellow struct {
							Zero15 struct {
								Total      int    `json:"total"`
								Percentage string `json:"percentage"`
							} `json:"0-15"`
							One630 struct {
								Total      int    `json:"total"`
								Percentage string `json:"percentage"`
							} `json:"16-30"`
							Three145 struct {
								Total      int    `json:"total"`
								Percentage string `json:"percentage"`
							} `json:"31-45"`
							Four660 struct {
								Total      int    `json:"total"`
								Percentage string `json:"percentage"`
							} `json:"46-60"`
							Six175 struct {
								Total      int    `json:"total"`
								Percentage string `json:"percentage"`
							} `json:"61-75"`
							Seven690 struct {
								Total      int    `json:"total"`
								Percentage string `json:"percentage"`
							} `json:"76-90"`
							Nine1105 struct {
								Total      interface{} `json:"total"`
								Percentage interface{} `json:"percentage"`
							} `json:"91-105"`
							One06120 struct {
								Total      interface{} `json:"total"`
								Percentage interface{} `json:"percentage"`
							} `json:"106-120"`
						} `json:"yellow"`
						Red struct {
							Zero15 struct {
								Total      interface{} `json:"total"`
								Percentage interface{} `json:"percentage"`
							} `json:"0-15"`
							One630 struct {
								Total      interface{} `json:"total"`
								Percentage interface{} `json:"percentage"`
							} `json:"16-30"`
							Three145 struct {
								Total      interface{} `json:"total"`
								Percentage interface{} `json:"percentage"`
							} `json:"31-45"`
							Four660 struct {
								Total      int    `json:"total"`
								Percentage string `json:"percentage"`
							} `json:"46-60"`
							Six175 struct {
								Total      int    `json:"total"`
								Percentage string `json:"percentage"`
							} `json:"61-75"`
							Seven690 struct {
								Total      interface{} `json:"total"`
								Percentage interface{} `json:"percentage"`
							} `json:"76-90"`
							Nine1105 struct {
								Total      interface{} `json:"total"`
								Percentage interface{} `json:"percentage"`
							} `json:"91-105"`
							One06120 struct {
								Total      interface{} `json:"total"`
								Percentage interface{} `json:"percentage"`
							} `json:"106-120"`
						} `json:"red"`
					} `json:"cards"`
				} `json:"league"`
			} `json:"away"`
		} `json:"teams"`
		Comparison struct {
			Form struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"form"`
			Att struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"att"`
			Def struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"def"`
			PoissonDistribution struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"poisson_distribution"`
			H2H struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"h2h"`
			Goals struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"goals"`
			Total struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"total"`
		} `json:"comparison"`
		H2H []struct {
			Fixture struct {
				ID        int       `json:"id"`
				Referee   string    `json:"referee"`
				Timezone  string    `json:"timezone"`
				Date      time.Time `json:"date"`
				Timestamp int       `json:"timestamp"`
				Periods   struct {
					First  int `json:"first"`
					Second int `json:"second"`
				} `json:"periods"`
				Venue struct {
					ID   interface{} `json:"id"`
					Name string      `json:"name"`
					City interface{} `json:"city"`
				} `json:"venue"`
				Status struct {
					Long    string `json:"long"`
					Short   string `json:"short"`
					Elapsed int    `json:"elapsed"`
				} `json:"status"`
			} `json:"fixture"`
			League struct {
				ID      int    `json:"id"`
				Name    string `json:"name"`
				Country string `json:"country"`
				Logo    string `json:"logo"`
				Flag    string `json:"flag"`
				Season  int    `json:"season"`
				Round   string `json:"round"`
			} `json:"league"`
			Teams struct {
				Home struct {
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Logo   string `json:"logo"`
					Winner bool   `json:"winner"`
				} `json:"home"`
				Away struct {
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Logo   string `json:"logo"`
					Winner bool   `json:"winner"`
				} `json:"away"`
			} `json:"teams"`
			Goals struct {
				Home int `json:"home"`
				Away int `json:"away"`
			} `json:"goals"`
			Score struct {
				Halftime struct {
					Home int `json:"home"`
					Away int `json:"away"`
				} `json:"halftime"`
				Fulltime struct {
					Home int `json:"home"`
					Away int `json:"away"`
				} `json:"fulltime"`
				Extratime struct {
					Home interface{} `json:"home"`
					Away interface{} `json:"away"`
				} `json:"extratime"`
				Penalty struct {
					Home interface{} `json:"home"`
					Away interface{} `json:"away"`
				} `json:"penalty"`
			} `json:"score"`
		} `json:"h2h"`
	} `json:"response"`
}

type Prediction struct {
	gorm.Model
	FixtureID int
	Percent   struct {
		Home string
		Draw string
		Away string
	}
	Home struct {
		ID   int
		Name string
	}
	Away struct {
		ID   int
		Name string
	}
	Comparison struct {
		Form struct {
			Home string `json:"home"`
			Away string `json:"away"`
		} `json:"form"`
		Att struct {
			Home string `json:"home"`
			Away string `json:"away"`
		} `json:"att"`
		Def struct {
			Home string `json:"home"`
			Away string `json:"away"`
		} `json:"def"`
		PoissonDistribution struct {
			Home string `json:"home"`
			Away string `json:"away"`
		} `json:"poisson_distribution"`
		H2H struct {
			Home string `json:"home"`
			Away string `json:"away"`
		} `json:"h2h"`
		Goals struct {
			Home string `json:"home"`
			Away string `json:"away"`
		} `json:"goals"`
		Total struct {
			Home string `json:"home"`
			Away string `json:"away"`
		} `json:"total"`
	} `json:"comparison"`
}
