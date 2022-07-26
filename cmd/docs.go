package main

import (
	"fmt"
	"path/filepath"

	"emperror.dev/errors"
	"github.com/MakeNowJust/heredoc"
	"github.com/banzaicloud/operator-tools/pkg/docgen"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var logger = zap.New(zap.UseDevMode(true))

func main() {
	// plugins()
	crds()
}

func plugins() {
	lister := docgen.NewSourceLister(
		map[string]docgen.SourceDir{
			"filters": {Path: "pkg/sdk/logging/model/filter", DestPath: "docs/configuration/plugins/filters"},
			"outputs": {Path: "pkg/sdk/logging/model/output", DestPath: "docs/configuration/plugins/outputs"},
			"common":  {Path: "pkg/sdk/logging/model/common", DestPath: "docs/configuration/plugins/common"},
		},
		logger.WithName("pluginlister"))

	lister.IgnoredSources = []string{
		"null",
		".*.deepcopy",
		".*_test",
	}

	lister.DefaultValueFromTagExtractor = func(tag string) string {
		return docgen.GetPrefixedValue(tag, `plugin:\"default:(.*)\"`)
	}

	lister.Index = docgen.NewDoc(docgen.DocItem{
		Name:     "_index",
		DestPath: "docs/configuration/plugins",
	}, logger.WithName("plugins"))

	lister.Header = heredoc.Doc(`
		---
		title: Supported Plugins
		generated_file: true
		---
		# Supported Plugins
		
		For more information please click on the plugin name
		<center>
		| Name | Profile | Description | Status |Version |
		|:---|---|:---|:---:|---:|`,
	)

	lister.Footer = heredoc.Doc(`
		</center>
	`)

	lister.DocGeneratedHook = func(document *docgen.Doc) error {
		relPath, err := filepath.Rel(lister.Index.Item.DestPath, document.Item.DestPath)
		if err != nil {
			return errors.WrapIff(err, "failed to determine relpath for %s", document.Item.DestPath)
		}

		lister.Index.Append(fmt.Sprintf("| **[%s](%s/)** | %s | %s | %s | [%s](%s) |",
			document.DisplayName,
			filepath.Join(relPath, document.Item.Name),
			document.Item.Category,
			document.Desc,
			document.Status,
			document.Version,
			document.Url))
		return nil
	}

	if err := lister.Generate(); err != nil {
		panic(err)
	}
}

func crds() {
	lister := docgen.NewSourceLister(
		map[string]docgen.SourceDir{
			"v1alpha1": {Path: "api/v1alpha1", DestPath: "docs/configuration/crds/v1alpha1"},
		},
		logger.WithName("crdlister"))

	lister.IgnoredSources = []string{
		".*.deepcopy",
		".*_test",
		".*_info",
	}

	lister.DefaultValueFromTagExtractor = func(tag string) string {
		return docgen.GetPrefixedValue(tag, `plugin:\"default:(.*)\"`)
	}

	lister.Index = docgen.NewDoc(docgen.DocItem{
		Name:     "_index",
		DestPath: "docs/configuration/crds/v1alpha1",
	}, logger.WithName("crds"))

	lister.Header = heredoc.Doc(`
		---
		title: Available CRDs
		generated_file: true
		---
	
		For more information please click on the name
		<center>
		| Name | Description | Version |
		|---|---|---|`,
	)

	lister.Footer = heredoc.Doc(`
		</center>
	`)

	lister.DocGeneratedHook = func(document *docgen.Doc) error {
		relPath, err := filepath.Rel(lister.Index.Item.DestPath, document.Item.DestPath)
		if err != nil {
			return errors.WrapIff(err, "failed to determine relpath for %s", document.Item.DestPath)
		}
		lister.Index.Append(fmt.Sprintf("| **[%s](%s/)** | %s | %s |",
			document.DisplayName,
			filepath.Join(relPath, document.Item.Name),
			document.Desc,
			document.Item.Category))
		return nil
	}

	if err := lister.Generate(); err != nil {
		panic(err)
	}
}
