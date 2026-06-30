package action

import "github.com/carapace-sh/carapace"

// ActionSeqtypes completes the sequence data types accepted by --seqtype.
func ActionSeqtypes() carapace.Action {
	return carapace.ActionValues(
		"AA",
		"BIN",
		"CODON",
		"DNA",
		"MORPH",
		"NT2AA",
	).Usage("BIN, DNA, AA, NT2AA, CODON, MORPH (auto-detected when omitted)")
}

// ActionSamplings completes the resampling strategies for --sampling.
func ActionSamplings() carapace.Action {
	return carapace.ActionValuesDescribed(
		"SITE", "resample sites",
		"GENE", "resample partition (gene)",
		"GENESITE", "resample both gene and sites",
	).Tag("sampling")
}

// ActionMergeAlgorithms completes the partition merging algorithms for --merge.
func ActionMergeAlgorithms() carapace.Action {
	return carapace.ActionValuesDescribed(
		"greedy", "greedy merging",
		"rcluster", "relaxed clustering algorithm",
		"rclusterf", "fast relaxed clustering algorithm",
	).Tag("merge algorithms")
}

// ActionMergeScopes completes the 1|all scope for --merge-model and --merge-rate.
func ActionMergeScopes() carapace.Action {
	return carapace.ActionValues(
		"1",
		"all",
	).Tag("merge scopes")
}

// ActionMerits completes the information criteria for --merit.
func ActionMerits() carapace.Action {
	return carapace.ActionValuesDescribed(
		"AIC", "Akaike information criterion",
		"AICc", "Akaike information criterion (corrected)",
		"BIC", "Bayesian information criterion",
	).Tag("information criteria")
}

// ActionMsets completes the program model sets for --mset.
func ActionMsets() carapace.Action {
	return carapace.ActionValues(
		"raxml",
		"phyml",
		"mrbayes",
		"beast1",
		"beast2",
	).Tag("program model sets")
}

// ActionMsubs completes the amino-acid model sources for --msub.
func ActionMsubs() carapace.Action {
	return carapace.ActionValues(
		"nuclear",
		"mitochondrial",
		"chloroplast",
		"viral",
	).Tag("amino-acid model sources")
}

// ActionSymtestTypes completes the symmetry test types for --symtest-type.
func ActionSymtestTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"MAR", "marginal test",
		"INT", "internal test",
	).Tag("symmetry test types")
}

// ActionRandTrees completes the random tree generation schemes for --rand.
func ActionRandTrees() carapace.Action {
	return carapace.ActionValuesDescribed(
		"UNI", "uniform random tree",
		"CAT", "caterpillar tree",
		"BAL", "balanced tree",
	).Tag("random tree schemes")
}

// ActionDatingMethods completes the dating methods for --dating.
func ActionDatingMethods() carapace.Action {
	return carapace.ActionValuesDescribed(
		"LSD", "least squares dating",
	).Tag("dating methods")
}

// ActionAfFormats completes the AliSim output formats for -af.
func ActionAfFormats() carapace.Action {
	return carapace.ActionValues(
		"phy",
		"fasta",
	).Tag("alignment formats")
}

// ActionSiteFreqOptions completes the MEAN|SAMPLING|MODEL options used by
// AliSim's --site-freq and --site-rate.
func ActionSiteFreqOptions() carapace.Action {
	return carapace.ActionValues(
		"MEAN",
		"SAMPLING",
		"MODEL",
	).Tag("site frequency options")
}

// ActionModels completes base substitution model names and the ModelFinder
// keywords for -m. Models are composed further with state-frequency (+F, +FO,
// ...) and rate-heterogeneity (+I, +G[n], +R[n], +Hn, ...) suffixes.
func ActionModels() carapace.Action {
	return carapace.ActionValuesDescribed(
		// ModelFinder keywords
		"TEST", "model selection followed by tree inference",
		"TESTONLY", "standard model selection only",
		"MF", "extended model selection with FreeRate heterogeneity",
		"MFP", "extended model selection followed by tree inference",
		// DNA
		"JC", "DNA: Jukes-Cantor",
		"K2P", "DNA: Kimura 2-parameter",
		"F81", "DNA: Felsenstein 1981",
		"HKY", "DNA: Hasegawa-Kishino-Yano",
		"K3P", "DNA: Kimura 3-parameter",
		"K81uf", "DNA: Kimura 1981 unequal frequencies",
		"TN", "DNA: Tam-Nei",
		"TrN", "DNA: Tam-Nei (alias)",
		"TNef", "DNA: Tam-Nei equal frequencies",
		"TIM", "DNA: transition model",
		"TIMef", "DNA: transition model equal frequencies",
		"TVM", "DNA: transversion model",
		"TVMef", "DNA: transversion model equal frequencies",
		"SYM", "DNA: symmetrical model",
		"GTR", "DNA: general time reversible",
		// Protein
		"Dayhoff", "protein: Dayhoff",
		"DCMut", "protein: DCMut",
		"FLAVI", "protein: FLAVI",
		"FLU", "protein: FLU",
		"HIVb", "protein: HIV between",
		"HIVw", "protein: HIV within",
		"JTT", "protein: Jones-Taylor-Thornton",
		"JTTDCMut", "protein: JTT-DCMut",
		"LG", "protein: Le-Gascuel",
		"mtART", "protein: mtART",
		"mtMAM", "protein: mtMAM",
		"mtMet", "protein: mtMet",
		"mtREV", "protein: mtREV",
		"mtInv", "protein: mtInv",
		"mtVer", "protein: mtVer",
		"mtZOA", "protein: mtZOA",
		"PMB", "protein: PMB",
		"Poisson", "protein: Poisson",
		"Q.bird", "protein: Q.bird",
		"Q.insect", "protein: Q.insect",
		"Q.LG", "protein: Q.LG",
		"Q.mammal", "protein: Q.mammal",
		"Q.pfam", "protein: Q.pfam",
		"Q.pfam_gb", "protein: Q.pfam_gb",
		"Q.plant", "protein: Q.plant",
		"Q.yeast", "protein: Q.yeast",
		"rtREV", "protein: rtREV",
		"VT", "protein: VT",
		"WAG", "protein: WAG",
		"GTR20", "protein: GTR20",
		"Blosum62", "protein: Blosum62",
		"cpREV", "protein: cpREV",
		// Protein mixture
		"C10", "protein mixture: C10",
		"C20", "protein mixture: C20",
		"C30", "protein mixture: C30",
		"C40", "protein mixture: C40",
		"C50", "protein mixture: C50",
		"C60", "protein mixture: C60",
		"EHO", "protein mixture: EHO",
		"EX2", "protein mixture: EX2",
		"EX3", "protein mixture: EX3",
		"EX_EHO", "protein mixture: EX_EHO",
		"LG4M", "protein mixture: LG4M",
		"LG4X", "protein mixture: LG4X",
		"UL2", "protein mixture: UL2",
		"UL3", "protein mixture: UL3",
		// Binary
		"JC2", "binary: JC2",
		"GTR2", "binary: GTR2",
		// Empirical codon
		"KOSI07", "empirical codon: KOSI07",
		"SCHN05", "empirical codon: SCHN05",
		// Mechanistic codon
		"GY", "mechanistic codon: Goldman-Yang",
		"GY0K", "mechanistic codon: GY0K",
		"GY1KTS", "mechanistic codon: GY1KTS",
		"GY1KTV", "mechanistic codon: GY1KTV",
		"GY2K", "mechanistic codon: GY2K",
		"MG", "mechanistic codon: Muse-Gaut",
		"MGK", "mechanistic codon: MGK",
		"MG1KTS", "mechanistic codon: MG1KTS",
		"MG1KTV", "mechanistic codon: MG1KTV",
		"MG2K", "mechanistic codon: MG2K",
		// Morphology / SNP
		"MK", "morphology/SNP: MK",
		"ORDERED", "morphology/SNP: ORDERED",
		// Lie Markov DNA
		"1.1", "Lie Markov DNA",
		"2.2b", "Lie Markov DNA",
		"3.3a", "Lie Markov DNA",
		"3.3b", "Lie Markov DNA",
		"3.3c", "Lie Markov DNA",
		"3.4", "Lie Markov DNA",
		"4.4a", "Lie Markov DNA",
		"4.4b", "Lie Markov DNA",
		"4.5a", "Lie Markov DNA",
		"4.5b", "Lie Markov DNA",
		"5.6a", "Lie Markov DNA",
		"5.6b", "Lie Markov DNA",
		"5.7a", "Lie Markov DNA",
		"5.7b", "Lie Markov DNA",
		"5.7c", "Lie Markov DNA",
		"5.11a", "Lie Markov DNA",
		"5.11b", "Lie Markov DNA",
		"5.11c", "Lie Markov DNA",
		"5.16", "Lie Markov DNA",
		"6.6", "Lie Markov DNA",
		"6.7a", "Lie Markov DNA",
		"6.7b", "Lie Markov DNA",
		"6.8a", "Lie Markov DNA",
		"6.8b", "Lie Markov DNA",
		"6.17a", "Lie Markov DNA",
		"6.17b", "Lie Markov DNA",
		"8.8", "Lie Markov DNA",
		"8.10a", "Lie Markov DNA",
		"8.10b", "Lie Markov DNA",
		"8.16", "Lie Markov DNA",
		"8.17", "Lie Markov DNA",
		"8.18", "Lie Markov DNA",
		"9.20a", "Lie Markov DNA",
		"9.20b", "Lie Markov DNA",
		"10.12", "Lie Markov DNA",
		"10.34", "Lie Markov DNA",
		"12.12", "Lie Markov DNA",
		// Non-reversible
		"NONREV", "non-reversible model",
		"UNREST", "unrestricted model (equiv. 12.12)",
		"STRSYM", "strand-symmetric model (equiv. WS6.6)",
		"NQ.bird", "non-reversible Q.bird",
		"NQ.insect", "non-reversible Q.insect",
		"NQ.mammal", "non-reversible Q.mammal",
		"NQ.pfam", "non-reversible Q.pfam",
		"NQ.plant", "non-reversible Q.plant",
		"NQ.yeast", "non-reversible Q.yeast",
	).Tag("substitution models")
}
