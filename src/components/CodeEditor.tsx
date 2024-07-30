import React, { useEffect, useState } from 'react';
import CodeMirror from '@uiw/react-codemirror';
import { dracula } from '@uiw/codemirror-theme-dracula';
import { ViewUpdate, lineNumbers } from '@codemirror/view';
import { languages } from '../languages';

type CodeEditorProps = {
	language: string;
	code: string,
	setCode: React.Dispatch<React.SetStateAction<string>>
}

export const CodeEditor: React.FC<CodeEditorProps> = ({language, code, setCode}) => {
    const [extensions, setExtensions] = useState<any[]>([lineNumbers()]);

	const onChange = React.useCallback((value: string, _: ViewUpdate) => {
		setCode(value);
	}, []);

	useEffect(() => {
		const loadLanguage = async () => {
			const lang = await languages[language as keyof typeof languages]();
			setExtensions([lineNumbers(), lang.languageSupport()]);
		}
		loadLanguage();
	}, [language]);

	return (
		<CodeMirror
			value={code}
			height="75vh"
			theme={dracula}
			extensions={extensions}
			onChange={onChange}
		/>
	);
};
