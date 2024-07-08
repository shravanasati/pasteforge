import React, { useState } from 'react';
import CodeMirror from '@uiw/react-codemirror';
import { cpp } from '@codemirror/lang-cpp';
import { dracula } from '@uiw/codemirror-theme-dracula';
import { ViewUpdate } from '@codemirror/view';

export const CodeEditor: React.FC = () => {
	const [code, setCode] = useState<string>('// Type your C++ code here');

	const onChange = React.useCallback((value: string, _: ViewUpdate) => {
		setCode(value);
	}, []);

	return (
		<CodeMirror
			value={code}
			height="200px"
			theme={dracula}
			extensions={[cpp()]}
			onChange={onChange}
		/>
	);
};
