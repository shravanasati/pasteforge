import React, { useState } from 'react';
import CodeMirror from '@uiw/react-codemirror';
import { dracula } from '@uiw/codemirror-theme-dracula';
import { ViewUpdate, lineNumbers } from '@codemirror/view';
import { languages } from '../languages';

type CodeEditorProps = {
	language: string;
}

export const CodeEditor: React.FC<CodeEditorProps> = ({language}) => {
	const [code, setCode] = useState<string>('');

	const onChange = React.useCallback((value: string, _: ViewUpdate) => {
		setCode(value);
	}, []);

	const langExt = languages[language as keyof typeof languages];

	const exts = [lineNumbers()];
	if (langExt) {
		exts.push(langExt());
	}

	return (
		<CodeMirror
			value={code}
			height="75vh"
			theme={dracula}
			extensions={exts}
			onChange={onChange}
		/>
	);
};
