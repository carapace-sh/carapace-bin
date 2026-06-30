package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/iqtree3_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "iqtree3",
	Short: "efficient phylogenomic software by maximum likelihood",
	Long:  "https://iqtree.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("Q", "Q", "", "Like -p but edge-unlinked partition model")
	rootCmd.Flags().StringP("S", "S", "", "Like -p but separate tree inference")
	rootCmd.Flags().StringP("T", "T", "", "No. cores/threads or AUTO-detect (default: 1)")
	rootCmd.Flags().Bool("abayes", false, "Approximate Bayes test (Anisimova et al. 2011)")
	rootCmd.Flags().String("af", "", "AliSim output format (default: phylip)")
	rootCmd.Flags().String("alisim", "", "Activate AliSim and specify the output alignment filename")
	rootCmd.Flags().Bool("allnni", false, "Perform more thorough NNI search (default: OFF)")
	rootCmd.Flags().Bool("alninfo", false, "Print alignment sites statistics to .alninfo")
	rootCmd.Flags().String("alpha-min", "", "Min Gamma shape parameter for site rates (default: 0.02)")
	rootCmd.Flags().String("alrt", "", "Replicates for SH approximate likelihood ratio test (0 for parametric aLRT)")
	rootCmd.Flags().Bool("ancestral", false, "Ancestral state reconstruction by empirical Bayes")
	rootCmd.Flags().String("asr-min", "", "Min probability of ancestral state (default: equil freq)")
	rootCmd.Flags().String("au-epsilon", "", "Epsilon for AU test (default: 0.001)")
	rootCmd.Flags().String("bcon", "", "Replicates for bootstrap + consensus tree")
	rootCmd.Flags().String("bcor", "", "Minimum correlation coefficient (default: 0.99)")
	rootCmd.Flags().String("beps", "", "RELL epsilon to break tie (default: 0.5)")
	rootCmd.Flags().Bool("blfix", false, "Fix branch lengths of user tree passed via -te")
	rootCmd.Flags().String("blmax", "", "Max branch length for optimization (default 100)")
	rootCmd.Flags().String("blmin", "", "Min branch length for optimization (default 0.000001)")
	rootCmd.Flags().Bool("blscale", false, "Scale branch lengths of user tree passed via -t")
	rootCmd.Flags().Bool("bnni", false, "Optimize UFBoot trees by NNI on bootstrap alignment")
	rootCmd.Flags().String("bonly", "", "Replicates for bootstrap only")
	rootCmd.Flags().StringP("boot", "b", "", "Replicates for bootstrap + ML tree + consensus tree")
	rootCmd.Flags().Bool("boot-trees", false, "Write bootstrap trees to .ufboot file (default: none)")
	rootCmd.Flags().String("branch-distribution", "", "Distribution to randomly generate and override branch lengths")
	rootCmd.Flags().String("branch-scale", "", "Value to scale all branch lengths")
	rootCmd.Flags().String("burnin", "", "Burnin number of trees to ignore")
	rootCmd.Flags().Bool("cf-quartet", false, "Write sCF for all resampled quartets to .cf.quartet")
	rootCmd.Flags().Bool("cf-verbose", false, "Write CF per tree/locus to cf.stat_tree/_loci")
	rootCmd.Flags().String("clock-sd", "", "Std-dev for lognormal relaxed clock (default: 0.2)")
	rootCmd.Flags().String("cmax", "", "Max categories for FreeRate model [+R] (default: 10)")
	rootCmd.Flags().String("cmin", "", "Min categories for FreeRate model [+R] (default: 2)")
	rootCmd.Flags().Bool("con-net", false, "Compute consensus network to .nex file")
	rootCmd.Flags().Bool("con-tree", false, "Compute consensus tree to .contree file")
	rootCmd.Flags().String("cptime", "", "Minimum checkpoint interval (default: 60 sec and adapt)")
	rootCmd.Flags().String("date", "", "File containing dates of tips/ancestral nodes, or TAXNAME to extract from taxon names")
	rootCmd.Flags().String("date-ci", "", "Number of replicates to compute confidence interval")
	rootCmd.Flags().Bool("date-no-outgroup", false, "Exclude outgroup from time tree")
	rootCmd.Flags().String("date-options", "", "Extra options passing directly to LSD2")
	rootCmd.Flags().String("date-outlier", "", "Z-score cutoff to remove outlier tips/nodes (e.g. 3)")
	rootCmd.Flags().String("date-root", "", "Root date as a real number or YYYY-MM-DD")
	rootCmd.Flags().String("date-tip", "", "Tip dates as a real number or YYYY-MM-DD")
	rootCmd.Flags().String("dating", "", "Dating method: LSD for least square dating (default)")
	rootCmd.Flags().Bool("df-tree", false, "Write discordant trees associated with gDF1")
	rootCmd.Flags().String("distribution", "", "Definition file of distributions for random model parameters")
	rootCmd.Flags().Bool("eigenlib", false, "Use Eigen3 library")
	rootCmd.Flags().String("epsilon", "", "Likelihood epsilon for parameter estimate (default 0.01)")
	rootCmd.Flags().Bool("fast", false, "Fast search to resemble FastTree")
	rootCmd.Flags().String("fconst", "", "Add constant patterns into alignment (N=no. states)")
	rootCmd.Flags().Bool("freq-max", false, "Posterior maximum instead of mean approximation")
	rootCmd.Flags().String("fundi", "", "List of taxa and Rho (Fundi weight) for FunDi model")
	rootCmd.Flags().StringP("g", "g", "", "Multifurcating topological constraint tree file")
	rootCmd.Flags().Bool("g_non_stop", false, "Turn off all stopping rules")
	rootCmd.Flags().Bool("g_print", false, "Write all generated species-trees (may be millions)")
	rootCmd.Flags().Bool("g_print_induced", false, "Write induced partition subtrees")
	rootCmd.Flags().String("g_print_lim", "", "Limit on the number of species-trees to be written")
	rootCmd.Flags().Bool("g_print_m", false, "Write presence-absence matrix")
	rootCmd.Flags().String("g_query", "", "Species-trees to test for identical set of subtrees")
	rootCmd.Flags().String("g_rm_leaves", "", "Invoke reverse analysis for complex datasets")
	rootCmd.Flags().String("g_stop_h", "", "Stop after NUM hours CPU time (0 to turn off, default: 7 days)")
	rootCmd.Flags().String("g_stop_i", "", "Stop after NUM intermediate trees (0 to turn off, default: 10MLN)")
	rootCmd.Flags().String("g_stop_t", "", "Stop after NUM species-trees (0 to turn off, default: 1MLN)")
	rootCmd.Flags().Bool("gamma-median", false, "Median approximation for +G site rates (default: mean)")
	rootCmd.Flags().String("gcf", "", "Set of source trees for gene concordance factor (gCF)")
	rootCmd.Flags().String("gentrius", "", "File with a single species-tree or a set of subtrees for Gentrius")
	rootCmd.Flags().Bool("gz", false, "Enable AliSim output compression (longer runtime)")
	rootCmd.Flags().BoolP("help", "h", false, "Print (more) help usages")
	rootCmd.Flags().String("indel", "", "Insertion and deletion rate of the indel model, relative to substitution rate")
	rootCmd.Flags().String("indel-size", "", "Insertion and deletion size distributions for the indel model")
	rootCmd.Flags().StringP("jack", "j", "", "Replicates for jackknife + ML tree + consensus tree")
	rootCmd.Flags().String("jack-prop", "", "Subsampling proportion for jackknife (default: 0.5)")
	rootCmd.Flags().Bool("keep-ident", false, "Keep identical sequences (default: remove & finally add)")
	rootCmd.Flags().String("lbp", "", "Replicates for fast local bootstrap probabilities")
	rootCmd.Flags().String("length", "", "Length of the root sequence (AliSim)")
	rootCmd.Flags().String("lmap", "", "Number of quartets for likelihood mapping analysis")
	rootCmd.Flags().String("lmclust", "", "NEXUS file containing clusters for likelihood mapping")
	rootCmd.Flags().StringP("m", "m", "", "Model name string (e.g. GTR+F+I+G)")
	rootCmd.Flags().String("madd", "", "Comma-separated list of mixture models to consider")
	rootCmd.Flags().String("mdef", "", "Model definition NEXUS file (see Manual)")
	rootCmd.Flags().String("mem", "", "Maximal RAM usage in GB | MB | %")
	rootCmd.Flags().String("merge", "", "Merge partitions to increase model fit (greedy|rcluster|rclusterf)")
	rootCmd.Flags().String("merge-model", "", "Use only 1 or all models for merging (default: 1)")
	rootCmd.Flags().String("merge-rate", "", "Use only 1 or all rate heterogeneity for merging (default: 1)")
	rootCmd.Flags().String("merit", "", "Akaike|Bayesian information criterion (default: BIC)")
	rootCmd.Flags().String("mfreq", "", "Comma-separated list of state frequencies")
	rootCmd.Flags().Bool("mix-opt", false, "Optimize mixture weights (default: detect)")
	rootCmd.Flags().Bool("mlrate", false, "Write maximum likelihood site rates to .mlrate file")
	rootCmd.Flags().Bool("modelomatic", false, "Find best codon/protein/DNA models (Whelan et al. 2015)")
	rootCmd.Flags().String("mrate", "", "Comma-separated list of rate heterogeneity among sites")
	rootCmd.Flags().Bool("mrbayes", false, "Output a MrBayes block file as a template for future analysis")
	rootCmd.Flags().String("mset", "", "Restrict search to models supported by other programs")
	rootCmd.Flags().String("msub", "", "Amino-acid model source")
	rootCmd.Flags().Bool("mtree", false, "Perform full tree search for every model")
	rootCmd.Flags().StringP("n", "n", "", "Fix number of iterations to stop (default: OFF)")
	rootCmd.Flags().String("nbest", "", "Number of best trees retained during search (default: 5)")
	rootCmd.Flags().String("ninit", "", "Number of initial parsimony trees (default: 100)")
	rootCmd.Flags().String("nmax", "", "Maximum number of iterations (default: 1000)")
	rootCmd.Flags().String("nn-path-model", "", "Neural network file for substitution model (onnx format)")
	rootCmd.Flags().String("nn-path-rates", "", "Neural network file for alpha value (onnx format)")
	rootCmd.Flags().Bool("no-copy-gaps", false, "Disable copying gaps from input alignment (default: false)")
	rootCmd.Flags().Bool("no-outfiles", false, "Suppress printing output files")
	rootCmd.Flags().Bool("no-unaligned", false, "Disable outputting unaligned sequences file with indel models")
	rootCmd.Flags().String("nstep", "", "Iterations for UFBoot stopping rule (default: 100)")
	rootCmd.Flags().String("nstop", "", "Number of unsuccessful iterations to stop (default: 100)")
	rootCmd.Flags().String("ntop", "", "Number of top initial trees (default: 20)")
	rootCmd.Flags().String("num-alignments", "", "Number of output datasets (AliSim)")
	rootCmd.Flags().StringP("o", "o", "", "Outgroup taxon (list) for writing .treefile")
	rootCmd.Flags().StringP("p", "p", "", "NEXUS/RAxML partition file or directory with alignments")
	rootCmd.Flags().Bool("partlh", false, "Write partition log-likelihoods to .partlh file")
	rootCmd.Flags().Bool("pathogen", false, "Apply CMAPLE tree search if sequence divergence is low, else IQ-TREE")
	rootCmd.Flags().Bool("pathogen-force", false, "Apply CMAPLE tree search regardless of sequence divergence")
	rootCmd.Flags().String("perturb", "", "Perturbation strength for randomized NNI (default: 0.5)")
	rootCmd.Flags().Bool("polytomy", false, "Collapse near-zero branches into polytomy")
	rootCmd.Flags().String("pr_ab_matrix", "", "Presence-absence matrix of loci coverage")
	rootCmd.Flags().String("prefix", "", "Prefix for all output files (default: aln/partition)")
	rootCmd.Flags().StringP("q", "q", "", "Like -p but edge-linked equal partition model")
	rootCmd.Flags().Bool("quartetlh", false, "Print quartet log-likelihoods to .quartetlh file")
	rootCmd.Flags().Bool("quiet", false, "Quiet mode, suppress printing to screen (stdout)")
	rootCmd.Flags().StringP("r", "r", "", "No. taxa for Yule-Harding random tree")
	rootCmd.Flags().String("radius", "", "Radius for parsimony SPR search (default: 6)")
	rootCmd.Flags().String("rand", "", "UNIform | CATerpillar | BALanced random tree")
	rootCmd.Flags().Bool("rate", false, "Write empirical Bayesian site rates to .rate file")
	rootCmd.Flags().String("rcluster", "", "Percentage of partition pairs for rcluster algorithm")
	rootCmd.Flags().String("rcluster-max", "", "Max number of partition pairs (default: 10*partitions)")
	rootCmd.Flags().String("rclusterf", "", "Percentage of partition pairs for rclusterf algorithm")
	rootCmd.Flags().Bool("redo", false, "Redo both ModelFinder and tree search")
	rootCmd.Flags().Bool("redo-tree", false, "Restore ModelFinder and only redo tree search")
	rootCmd.Flags().String("rlen", "", "min, mean and max random branch lengths")
	rootCmd.Flags().String("root-seq", "", "Specify the root sequence from an alignment (FILE,SEQ_NAME)")
	rootCmd.Flags().String("runs", "", "Number of independent runs (default: 1)")
	rootCmd.Flags().StringP("s", "s", "", "PHYLIP/FASTA/NEXUS/CLUSTAL/MSF alignment file(s) or directory")
	rootCmd.Flags().Bool("safe", false, "Safe likelihood kernel to avoid numerical underflow")
	rootCmd.Flags().String("sampling", "", "GENE|GENESITE resampling for partitions (default: SITE)")
	rootCmd.Flags().String("scf", "", "Number of quartets for site concordance factor (sCF)")
	rootCmd.Flags().String("scfl", "", "Like --scf but using likelihood (recommended)")
	rootCmd.Flags().String("seed", "", "Random seed number, normally used for debugging purpose")
	rootCmd.Flags().String("seqtype", "", "BIN, DNA, AA, NT2AA, CODON, MORPH (default: auto-detect)")
	rootCmd.Flags().Bool("show-lh", false, "Compute tree likelihood without optimisation")
	rootCmd.Flags().Bool("single-output", false, "Output all alignments into a single file (AliSim)")
	rootCmd.Flags().String("site-freq", "", "Input site frequency model file")
	rootCmd.Flags().String("site-rate", "", "Option to mimic discrete rate heterogeneity (MEAN|SAMPLING|MODEL)")
	rootCmd.Flags().Bool("sitelh", false, "Write site log-likelihoods to .sitelh file")
	rootCmd.Flags().Bool("sprta", false, "Compute SPRTA (DeMaio et al., 2024) branch supports")
	rootCmd.Flags().Bool("sprta-other-places", false, "Output alternative SPRs and their SPRTA supports")
	rootCmd.Flags().Bool("sprta-zero-branch", false, "Compute SPRTA supports for zero-length branches")
	rootCmd.Flags().Bool("sub-level-mixture", false, "Enable substitution-level mixture model simulation")
	rootCmd.Flags().String("subsample", "", "Randomly sub-sample partitions (negative for complement)")
	rootCmd.Flags().String("subsample-seed", "", "Random number seed for --subsample")
	rootCmd.Flags().String("sup-min", "", "Min split support, 0.5 for majority-rule consensus (default: 0)")
	rootCmd.Flags().String("support", "", "Assign support values into this tree from -t trees")
	rootCmd.Flags().String("suptag", "", "Node name (or ALL) to assign tree IDs where node occurs")
	rootCmd.Flags().Bool("symtest", false, "Perform three tests of symmetry")
	rootCmd.Flags().Bool("symtest-keep-zero", false, "Keep NAs in the tests")
	rootCmd.Flags().Bool("symtest-only", false, "Do --symtest then exit")
	rootCmd.Flags().String("symtest-pval", "", "P-value cutoff (default: 0.05)")
	rootCmd.Flags().Bool("symtest-remove-bad", false, "Do --symtest and remove bad partitions")
	rootCmd.Flags().Bool("symtest-remove-good", false, "Do --symtest and remove good partitions")
	rootCmd.Flags().String("symtest-type", "", "Use MARginal/INTernal test when removing partitions")
	rootCmd.Flags().StringP("t", "t", "", "Starting tree (FILE|PARS|RAND, default: 99 parsimony and BIONJ)")
	rootCmd.Flags().Bool("tbe", false, "Transfer bootstrap expectation")
	rootCmd.Flags().String("te", "", "Evaluate/fix the user tree")
	rootCmd.Flags().Bool("terrace", false, "Check if the tree lies on a phylogenetic terrace")
	rootCmd.Flags().String("test", "", "Replicates for topology test")
	rootCmd.Flags().Bool("test-au", false, "Approximately unbiased (AU) test (Shimodaira 2002)")
	rootCmd.Flags().Bool("test-weight", false, "Perform weighted KH and SH tests")
	rootCmd.Flags().String("threads-max", "", "Max number of threads for -T AUTO (default: all cores)")
	rootCmd.Flags().String("tree-dist", "", "Compute RF distances between -t trees and this set")
	rootCmd.Flags().Bool("tree-dist-all", false, "Compute all-to-all RF distances for -t trees")
	rootCmd.Flags().String("tree-dist2", "", "Compute RF distances between -t trees and this set (unequal taxon sets)")
	rootCmd.Flags().Bool("tree-fix", false, "Fix -t tree (no tree search performed)")
	rootCmd.Flags().String("tree-freq", "", "Input tree to infer site frequency model")
	rootCmd.Flags().Bool("treels", false, "Write locally optimal trees into .treels file")
	rootCmd.Flags().String("trees", "", "Set of trees to evaluate log-likelihoods")
	rootCmd.Flags().StringP("ufboot", "B", "", "Replicates for ultrafast bootstrap (>=1000)")
	rootCmd.Flags().StringP("ufjack", "J", "", "Replicates for ultrafast jackknife (>=1000)")
	rootCmd.Flags().Bool("undo", false, "Revoke finished run, used when changing some options")
	rootCmd.Flags().Bool("use-nn-model", false, "Use neural network for tree inference")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose mode, printing more messages to screen")
	rootCmd.Flags().BoolP("version", "V", false, "Display version number")
	rootCmd.Flags().Bool("wbtl", false, "Like --boot-trees but also writing branch lengths")
	rootCmd.Flags().Bool("write-all", false, "Enable outputting internal sequences (AliSim)")
	rootCmd.Flags().Bool("wslm", false, "Write site log-likelihoods per mixture class")
	rootCmd.Flags().Bool("wslmr", false, "Write site log-likelihoods per mixture+rate class")
	rootCmd.Flags().Bool("wslr", false, "Write site log-likelihoods per rate category")
	rootCmd.Flags().Bool("wspm", false, "Write site probabilities per mixture class")
	rootCmd.Flags().Bool("wspmr", false, "Write site probabilities per mixture+rate class")
	rootCmd.Flags().Bool("wspr", false, "Write site probabilities per rate category")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"Q":             carapace.ActionFiles(),
		"S":             carapace.ActionFiles(),
		"T":             carapace.ActionValues("AUTO"),
		"af":            action.ActionAfFormats(),
		"date":          carapace.ActionFiles(),
		"dating":        action.ActionDatingMethods(),
		"distribution":  carapace.ActionFiles(),
		"g":             carapace.ActionFiles(),
		"g_query":       carapace.ActionFiles(),
		"gcf":           carapace.ActionFiles(),
		"gentrius":      carapace.ActionFiles(),
		"lmclust":       carapace.ActionFiles(),
		"m":             action.ActionModels(),
		"mdef":          carapace.ActionFiles(),
		"merge":         action.ActionMergeAlgorithms(),
		"merge-model":   action.ActionMergeScopes(),
		"merge-rate":    action.ActionMergeScopes(),
		"merit":         action.ActionMerits(),
		"mset":          action.ActionMsets().UniqueList(","),
		"msub":          action.ActionMsubs(),
		"nn-path-model": carapace.ActionFiles(),
		"nn-path-rates": carapace.ActionFiles(),
		"p":             carapace.ActionFiles(),
		"pr_ab_matrix":  carapace.ActionFiles(),
		"q":             carapace.ActionFiles(),
		"rand":          action.ActionRandTrees(),
		"root-seq":      carapace.ActionFiles(),
		"s":             carapace.ActionFiles(),
		"sampling":      action.ActionSamplings(),
		"seqtype":       action.ActionSeqtypes(),
		"site-freq":     carapace.ActionFiles(),
		"site-rate":     action.ActionSiteFreqOptions(),
		"support":       carapace.ActionFiles(),
		"symtest-type":  action.ActionSymtestTypes(),
		"t":             carapace.ActionFiles(),
		"te":            carapace.ActionFiles(),
		"tree-dist":     carapace.ActionFiles(),
		"tree-dist2":    carapace.ActionFiles(),
		"tree-freq":     carapace.ActionFiles(),
		"trees":         carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
