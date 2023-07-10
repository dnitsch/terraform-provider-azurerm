package applications

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationArtifact struct {
	Name ApplicationArtifactName `json:"name"`
	Type ApplicationArtifactType `json:"type"`
	Uri  string                  `json:"uri"`
}
