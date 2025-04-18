package warlock

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

var ItemSetDeathmistRaiment = core.NewItemSet(core.ItemSet{
	Name: "Deathmist Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		// +8 All Resistances.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddResistances(8)
		},
		// 4pc: When struck in combat has a chance of causing the attacker to flee in terror for 2 seconds.
		// Increases damage and healing done by magical spells and effects by up to 23.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.SpellPower, 23)
		},
		// +200 Armor.
		8: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Armor, 200)
		},
	},
})

var ItemSetFelheartRaiment = core.NewItemSet(core.ItemSet{
	Name: "Felheart Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		// 3 pieces: Health or Mana gained from Drain Life and Drain Mana increased by 15%.
		// 5 pieces: Your pet gains 15 stamina and 100 spell resistance against all schools of magic.
		// 8 pieces: Mana cost of Shadow spells reduced by 15%.
		8: func(agent core.Agent) {
			warlock := agent.(WarlockAgent).GetWarlock()
			warlock.RegisterAura(core.Aura{
				Label: "Shadow Cost Reduction (Felheart Raiment)",
				OnInit: func(aura *core.Aura, sim *core.Simulation) {
					for _, spell := range warlock.Spellbook {
						if spell.SpellSchool.Matches(core.SpellSchoolShadow) && spell.Cost != nil {
							spell.Cost.Multiplier -= 15
						}
					}
				},
			})
		},
	},
})

var ItemSetNemesisRaiment = core.NewItemSet(core.ItemSet{
	Name: "Nemesis Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		// 3 pieces: Increases damage and healing done by magical spells and effects by up to 23.
		3: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.SpellPower, 23)
		},
		// 5 pieces: Your pet gains 20 stamina and 130 spell resistance against all schools of magic.
		// 8 pieces: Reduces the threat generated by your Destruction spells by 20%.
		8: func(agent core.Agent) {
			warlock := agent.(WarlockAgent).GetWarlock()
			warlock.RegisterAura(core.Aura{
				Label: "Decreased Destruction Threat (Nemesis Raiment)",
				OnInit: func(aura *core.Aura, sim *core.Simulation) {
					for _, spell := range warlock.Spellbook {
						if spell.Flags.Matches(WarlockFlagDestruction) {
							spell.ThreatMultiplier *= 0.80
						}
					}
				},
			})
		},
	},
})

var ItemSetDemoniacsThreads = core.NewItemSet(core.ItemSet{
	Name: "Demoniac's Threads",
	Bonuses: map[int32]core.ApplyEffect{
		// Increases damage and healing done by magical spells and effects by up to 12.
		2: func(agent core.Agent) {
			warlock := agent.(WarlockAgent).GetWarlock()
			warlock.AddStat(stats.SpellPower, 12)
		},
		// 3 pieces: Increases the damage of Corruption by 2%.
		3: func(agent core.Agent) {
			warlock := agent.(WarlockAgent).GetWarlock()
			warlock.RegisterAura(core.Aura{
				Label: "Improved Corruption (Demoniac's Threads)",
				OnInit: func(aura *core.Aura, sim *core.Simulation) {
					for _, spell := range warlock.Corruption {
						spell.DamageMultiplierAdditive += 0.02
					}
				},
			})
		},
		// 5 pieces: Decreases the cooldown of Death Coil by 15%.
		5: func(agent core.Agent) {
			warlock := agent.(WarlockAgent).GetWarlock()
			warlock.RegisterAura(core.Aura{
				Label: "Improved Death Coil (Demoniac's Threads)",
				OnInit: func(aura *core.Aura, sim *core.Simulation) {
					for _, spell := range warlock.DeathCoil {
						spell.CD.Duration = time.Duration(float64(spell.CD.Duration) * 0.85)
					}
				},
			})
		},
	},
})

var ItemSetDoomcallersAttire = core.NewItemSet(core.ItemSet{
	Name: "Doomcaller's Attire",
	Bonuses: map[int32]core.ApplyEffect{
		// 3 pieces: 5% increased damage on your Immolate spell.
		3: func(agent core.Agent) {
			warlock := agent.(WarlockAgent).GetWarlock()
			warlock.RegisterAura(core.Aura{
				Label: "Doomcaller Immolate Bonus",
				OnInit: func(aura *core.Aura, sim *core.Simulation) {
					for _, spell := range warlock.Immolate {
						spell.BaseDamageMultiplierAdditive += 0.05
					}
				},
			})
		},
		// 5 pieces: Reduces the mana cost of Shadow Bolt by 15%.
		5: func(agent core.Agent) {
			warlock := agent.(WarlockAgent).GetWarlock()
			warlock.RegisterAura(core.Aura{
				Label: "Doomcaller Reduced Shadow Bolt Cost",
				OnInit: func(aura *core.Aura, sim *core.Simulation) {
					for _, spell := range warlock.ShadowBolt {
						spell.Cost.Multiplier -= 15
					}
				},
			})
		},
	},
})

var ItemSetPlagueheartRaiment = core.NewItemSet(core.ItemSet{
	Name: "Plagueheart Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		// 2 pieces: Your Shadow Bolts now have a chance to heal you for 270 to 330.
		// 4 pieces: Increases damage caused by your Corruption by 12%.
		4: func(agent core.Agent) {
			warlock := agent.(WarlockAgent).GetWarlock()
			warlock.RegisterAura(core.Aura{
				Label: "Corruption (Plagueheart Raiment)",
				OnInit: func(aura *core.Aura, sim *core.Simulation) {
					for _, spell := range warlock.Corruption {
						spell.DamageMultiplierAdditive += 0.12
					}
				},
			})
		},
		// 6 pieces: Your spell critical hits generate 25% less threat. In addition, Corruption, Immolate, Curse of Agony, and Siphon Life generate 25% less threat.
		6: func(agent core.Agent) {
			warlock := agent.(WarlockAgent).GetWarlock()
			warlock.RegisterAura(core.Aura{
				Label: "Plagueheart",
				OnInit: func(aura *core.Aura, sim *core.Simulation) {
					for _, spell := range warlock.Corruption {
						spell.ThreatMultiplier *= 0.75
					}

					for _, spell := range warlock.Immolate {
						spell.ThreatMultiplier *= 0.75
					}

					for _, spell := range warlock.CurseOfAgony {
						spell.ThreatMultiplier *= 0.75
					}

					// TODO: Spell crit thread? Do we care?
				},
			})
		},
		// 8 pieces: Reduces health cost of your Life Tap by 12%.
	},
})
